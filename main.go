package main

import (
    "log"
    "net/http"

    "github.com/yourorg/claude-test/router"
    "github.com/yourorg/claude-test/db"
)

func main() {
    // Initialize database connection
    if err := db.Init(); err != nil {
        log.Fatalf("failed to initialize database: %v", err)
    }

    // Setup router
    r := router.NewRouter()

    // Start HTTP server
    addr := ":8080"
    log.Printf("starting server on %s", addr)
    if err := http.ListenAndServe(addr, r); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
