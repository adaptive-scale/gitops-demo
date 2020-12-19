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
    log.Println("Starting app for demo with deb")
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8090", nil))
}
