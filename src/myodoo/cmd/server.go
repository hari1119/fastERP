package cmd

import (
    "fmt"
    "net/http"
)

func RunServer() error {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "🧠 Hello from MyOdoo framework!")
    })

    fmt.Println("📡 Listening on http://localhost:8080")
    return http.ListenAndServe(":8080", nil)
}
