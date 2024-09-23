package hello

import (
    "net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World – GWS"))
}

func HelloWorldHTML(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Hello World – GWS</h1>"))
}

func HelloWorldJSON(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message":"Hello World - GWS"}`))
}
