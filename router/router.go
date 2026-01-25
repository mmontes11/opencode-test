package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/yourorg/claude-test/handler"
)

// NewRouter creates a new HTTP router with example routes.
func NewRouter() http.Handler {
    r := mux.NewRouter()

    // Simple health check endpoint
    r.HandleFunc("/health", handler.HealthCheck).Methods("GET")

    // Add more routes here as needed
    return r
}
