package main

import (
    "html/template"
    "net/http"
    "log"
)

type Page struct {
    Title   string
    Header  string
    Body    string
}

func index(w http.ResponseWriter, r *http.Request) {
    IndexPage := Page {
        Title: "Golang web apps",
        Header: "Welcome to golang",
        Body: "This is simple page",
    }

    t, err := template.ParseFiles("index.html")
    if err != nil {
        log.Print("Template error: ", err)
    }
    err = t.Execute(w, IndexPage)
    if err != nil {
        log.Print("Template executing error: ", err)
    }

}

func main() {
    http.HandleFunc("/", index)
    http.ListenAndServe(":8000", nil)
}
