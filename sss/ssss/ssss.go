package ssss

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

const VERSION = "0.0.1"

var (
	SSSSApp *App
	AppPath string
)

func init() {
	SSSSApp = NewApp(nil)
	AppPath, _ = os.Getwd()
}

type App struct {
	Handlers         *ControllerRegistor
	config           *Config
	StaticDirs       map[string]string
	TemplateRegistor *TemplateRegistor
}

func NewApp(config *Config) *App {
	cr := NewControllerRegistor()
	app := &App{Handlers: cr,
		config:           config,
		StaticDirs:       make(map[string]string),
		TemplateRegistor: NewTemplateRegistor()}
	cr.app = app
	return app
}

//method-http method, GET,POST,PUT,HEAD,DELETE,PATCH,OPTIONS,*
//path-URL path
//name - method on the container
func (app *App) Register(method string, path string, c IController, name string) *App {
	app.Handlers.Add(method, path, c, name)
	return app
}

func Register(method string, path string, c IController, name string) *App {
	SSSSApp.Register(method, path, c, name)
	return SSSSApp
}

func (app *App) SetStaticPath(url string, path string) *App {
	app.StaticDirs[url] = path
	return app
}

func SetStaticPath(url string, path string) *App {
	SSSSApp.StaticDirs[url] = path
	return SSSSApp
}

func AddTemplateExt(ext string) {
	SSSSApp.TemplateRegistor.AddTemplateExt(ext)
}

func AddFuncMap(key string, funname interface{}) error {
	return SSSSApp.TemplateRegistor.AddFuncMap(key, funname)
}

func (app *App) buildTemplate() {
	if app.config.TemplatePath != "" {
		app.TemplateRegistor.buildTemplate(app.config.TemplatePath)
	}
}

func (app *App) Run() {
	app.buildTemplate()
	if app.config.HttpAddr == "" {
		app.config.HttpAddr = "0.0.0.0"
	}
	addr := fmt.Sprintf("%s:%d", app.config.HttpAddr, app.config.HttpPort)
	var err error
	if app.config.UseFcgi {
		l, e := net.Listen("tcp", addr)
		if e != nil {
			log.Print("Listen: ", e)
		}
		err = fcgi.Serve(l, app.Handlers)
	} else {
		err = http.ListenAndServe(addr, app.Handlers)
	}
	if err != nil {
		log.Print("ListenAndServe: ", err)
		panic(err)
	}
}

func Run(config *Config) {
	SSSSApp.config = config
	SSSSApp.Run()
}
