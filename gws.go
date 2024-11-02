package main

import (
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "time"
    "gws/hello"
    "gws/syllabus"
    "gws/help"
)

func d6Handler(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    roll := rand.Intn(6) + 1
    fmt.Fprintf(w, "You rolled a: %d", roll)
}

func main() {
    http.HandleFunc("/hello-world", hello.HelloWorld)
    http.HandleFunc("/hello-world-html", hello.HelloWorldHTML)
    http.HandleFunc("/hello-world-json", hello.HelloWorldJSON)
    http.HandleFunc("/syllabus", syllabus.SyllabusHandler)
    http.HandleFunc("/help", help.HelpHandler)
    http.HandleFunc("/d6", d6Handler)

    fmt.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
