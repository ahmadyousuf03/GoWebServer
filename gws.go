package main

import (
    "fmt"
    "log"
    "net/http"
    "gws/hello"
    "gws/syllabus"
    "gws/help"
)

func main() {
    http.HandleFunc("/hello-world", hello.HelloWorld)
    http.HandleFunc("/hello-world-html", hello.HelloWorldHTML)
    http.HandleFunc("/hello-world-json", hello.HelloWorldJSON)
    http.HandleFunc("/syllabus", syllabus.SyllabusHandler)
    http.HandleFunc("/help", help.HelpHandler)

    fmt.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
