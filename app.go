package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println("api received")
    fmt.Fprintf(w, "Hi there, This is adaptive demo with env parameters "+os.Getenv("TEST_PARAM_1")+" no. 106")
}

func errHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("error=triggered Error triggered for adaptive demo 1")
    log.Println("error=demo for show advanced logging")
    fmt.Fprintf(w, "Error triggered.")
}

func main() {
    log.Println("Starting app for demo for debarshi")
    http.HandleFunc("/", handler)
    http.HandleFunc("/error", errHandler)

    log.Fatal(http.ListenAndServe(":80", nil))
}
