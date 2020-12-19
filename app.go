package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, This is adaptive demo no. 1")
}

func main() {
    log.Println("Starting app for demo 2")
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8090", nil))
}
