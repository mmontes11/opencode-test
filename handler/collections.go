package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mmontes11/opencode-test/db"
	"github.com/mmontes11/opencode-test/store"
)

// CollectionRequest represents the expected payload for creating or updating a collection.
// The fields mirror the database columns.
//
// Example JSON payload:
//
//	{"name": "My collection", "description": "A collection of items"}
type CollectionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ItemInCollectionRequest represents the payload for adding an item to a collection.
// Example: {"item_id": 42}
type ItemInCollectionRequest struct {
	ItemID int64 `json:"item_id"`
}

// CreateCollectionHandler handles POST /collections.
func CreateCollectionHandler(w http.ResponseWriter, r *http.Request) {
	var req CollectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	col, err := store.CreateCollection(db.DB, req.Name, req.Description)
	if err != nil {
		fmt.Printf("store.CreateCollection error: %v\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(col)
}

// GetCollectionHandler handles GET /collections/{id}.
func GetCollectionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	col, err := store.GetCollection(db.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "collection not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(col)
}

// ListCollectionHandler handles GET /collections.
func ListCollectionHandler(w http.ResponseWriter, r *http.Request) {
	cols, err := store.ListCollections(db.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cols)
}

// UpdateCollectionHandler handles PUT /collections/{id}.
func UpdateCollectionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var req CollectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	col, err := store.UpdateCollection(db.DB, id, req.Name, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(col)
}

// DeleteCollectionHandler handles DELETE /collections/{id}.
func DeleteCollectionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := store.DeleteCollection(db.DB, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// AddItemToCollectionHandler handles POST /collections/{id}/items.
func AddItemToCollectionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	colID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid collection id", http.StatusBadRequest)
		return
	}
	var req ItemInCollectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.ItemID == 0 {
		http.Error(w, "item_id is required", http.StatusBadRequest)
		return
	}
	if err := store.AddItemToCollection(db.DB, colID, req.ItemID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ListItemsInCollectionHandler handles GET /collections/{id}/items.
func ListItemsInCollectionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	colID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid collection id", http.StatusBadRequest)
		return
	}
	items, err := store.ListItemsInCollection(db.DB, colID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// RemoveItemFromCollectionHandler handles DELETE /collections/{id}/items/{item_id}.
func RemoveItemFromCollectionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	colID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid collection id", http.StatusBadRequest)
		return
	}
	itemID, err := strconv.ParseInt(vars["item_id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid item id", http.StatusBadRequest)
		return
	}
	if err := store.RemoveItemFromCollection(db.DB, colID, itemID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
