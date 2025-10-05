package main

import (
    "fmt"
    "log"
    "net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Go service is up and running!")
}

func main() {
    http.HandleFunc("/health", healthHandler)

    fmt.Println("Go server is starting on port 8081...")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal(err)
    }
}