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

	// Collection routes
	r.HandleFunc("/collections", handler.CreateCollectionHandler).Methods("POST")
	r.HandleFunc("/collections", handler.ListCollectionHandler).Methods("GET")
	r.HandleFunc("/collections/{id}", handler.GetCollectionHandler).Methods("GET")
	r.HandleFunc("/collections/{id}", handler.UpdateCollectionHandler).Methods("PUT")
	r.HandleFunc("/collections/{id}", handler.DeleteCollectionHandler).Methods("DELETE")

	// Collection item routes
	r.HandleFunc("/collections/{id}/items", handler.AddItemToCollectionHandler).Methods("POST")
	r.HandleFunc("/collections/{id}/items", handler.ListItemsInCollectionHandler).Methods("GET")
	r.HandleFunc("/collections/{id}/items/{item_id}", handler.RemoveItemFromCollectionHandler).Methods("DELETE")

	return r
}
