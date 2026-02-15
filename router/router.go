package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmontes11/opencode-test/handler"
)

// NewRouter creates a new HTTP router with example routes.
func NewRouter() http.Handler {
	r := mux.NewRouter()

	// Simple health check endpoint
	r.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	// CRUD routes for items
	r.HandleFunc("/items", handler.CreateItemHandler).Methods("POST")
	r.HandleFunc("/items", handler.ListItemHandler).Methods("GET")
	r.HandleFunc("/items/{id}", handler.GetItemHandler).Methods("GET")
	r.HandleFunc("/items/{id}", handler.UpdateItemHandler).Methods("PUT")
	r.HandleFunc("/items/{id}", handler.DeleteItemHandler).Methods("DELETE")

	return r
}
