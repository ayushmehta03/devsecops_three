
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go DevSecOps Lab!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Application is healthy")
}

func newRouter() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)
    mux.HandleFunc("/health", healthHandler)
    return mux
}

func main() {
    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", newRouter()))
}
