package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, This is adaptive demo no. 8")
}

func main() {
    log.Println("Starting app for demo for debarshi")
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":80", nil))
}
