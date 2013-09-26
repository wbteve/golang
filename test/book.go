package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "time"
    //"os"
)

func main() {
    for {
        run()
        time.Sleep(time.Minute)
    }
}

func run() {
    dest := "http://read.douban.com/ebook/393187/"
    content := ""
    response, err := http.Get(dest)
    if err != nil {
        //handle error
        fmt.Println(err)
        log.Fatal(err)
    }
    if response.StatusCode == http.StatusOK {
        fmt.Println(time.Now().String() + "      " + string(response.StatusCode))
    }
    defer response.Body.Close()

    buf := make([]byte, 1024)

    for {
        n, _ := response.Body.Read(buf)
        if 0 == n {
            break
        }
        content += (string(buf[:n]))
    }
    strDict := strings.Split(content, "reader-actions")
    piece := strings.Split(strDict[1], "ratings-count")
    userString := piece[0]
    //fmt.Println(userString)
    usersDict := strings.Split(userString, "http://www.douban.com/people/")
    //fmt.Println(usersDict)
    for i := 0; i < len(usersDict); i++ {
        if i != 0 {
            if i%2 == 0 {
                uid, info := parseItem(usersDict[i])
                fmt.Println(uid + " " + info)
            }
        }
    }
}

func parseItem(userRecord string) (string, string) {
    s := strings.Split(userRecord, "/\"")
    remain := strings.Split(s[1], "action-msg\"><span>")[1]
    c := strings.Split(remain, "</span>")[0]
    info := strings.Split(c, " ")[1]
    return s[0], info
}
