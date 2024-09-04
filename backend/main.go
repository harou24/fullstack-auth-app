// main.go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/secure-data", secureEndpoint).Methods("GET")
    http.ListenAndServe(":8081", r)
}

func secureEndpoint(w http.ResponseWriter, r *http.Request) {
    // JWT validation and extraction logic
    w.Write([]byte("This is a secure endpoint!"))
}

