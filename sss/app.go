package main

import (
    "./ssss"
)

type MainController struct {
    ssss.Controller
}

func (this *MainController) Hello() {
    this.RenderText("hello world")
}

func (this *MainController) Hi() {
    this.RenderHtml("<html><body>Hi,<a href='/'>back</a></body></html>")
}

func (this *MainController) Htmlfile() {
	this.TplNames = "hi.html"
	this.RenderTemplate()
}

func main() {
    ssss.Register("GET|POST", "/", &MainController{}, "Htmlfile")
    ssss.Register("GET|POST", "/hi", &MainController{}, "Hi")
    ssss.Register("GET|POST", "/hello", &MainController{}, "Hello")
    
	ssss.SetStaticPath("/", "webroot/")

    var cfg ssss.Config
    //cfg.HttpAddr = "0.0.0.0"
    cfg.HttpPort = 8080
    
	cfg.TemplatePath = "templates"
	cfg.RunMode = 1 
    ssss.Run(&cfg)
}
