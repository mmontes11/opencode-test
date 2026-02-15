package handler

import (
	"bytes"
	"github.com/mmontes11/opencode-test/db"
	"github.com/mmontes11/opencode-test/store"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Helper to initialize DB or skip test if DSN not set.
func initDBForTest(t *testing.T) {
	dsn := os.Getenv("MARIADB_DSN")
	if dsn == "" {
		t.Skip("MARIADB_DSN env var not set; skipping integration tests")
	}
	if err := db.Init(); err != nil {
		t.Fatalf("failed to init db: %v", err)
	}
}

func TestCreateCollection(t *testing.T) {
	initDBForTest(t)
	// Clean up after
	defer db.DB.Close()
	col, err := store.CreateCollection(db.DB, "test collection", "test description")
	if err != nil {
		t.Fatalf("CreateCollection returned error: %v", err)
	}
	if col.ID == 0 || col.Name != "test collection" {
		t.Fatalf("unexpected collection data: %+v", col)
	}
	// delete
	if err := store.DeleteCollection(db.DB, col.ID); err != nil {
		t.Fatalf("DeleteCollection returned error: %v", err)
	}
}

func TestAddAndListItemsInCollection(t *testing.T) {
	initDBForTest(t)
	defer db.DB.Close()

	// create collection
	col, err := store.CreateCollection(db.DB, "col with items", "")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	// create items
	item1, _ := store.CreateItem(db.DB, "item1", "")
	item2, _ := store.CreateItem(db.DB, "item2", "")
	// add to collection
	if err := store.AddItemToCollection(db.DB, col.ID, item1.ID); err != nil {
		t.Fatalf("add failed: %v", err)
	}
	if err := store.AddItemToCollection(db.DB, col.ID, item2.ID); err != nil {
		t.Fatalf("add failed: %v", err)
	}
	// list
	items, err := store.ListItemsInCollection(db.DB, col.ID)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	// cleanup
	store.DeleteItem(db.DB, item1.ID)
	store.DeleteItem(db.DB, item2.ID)
	store.DeleteCollection(db.DB, col.ID)
}

func TestDeleteCollection(t *testing.T) {
	initDBForTest(t)
	defer db.DB.Close()
	col, err := store.CreateCollection(db.DB, "to delete", "")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if err := store.DeleteCollection(db.DB, col.ID); err != nil {
		t.Fatalf("err: %v", err)
	}
	// attempt get
	if _, err := store.GetCollection(db.DB, col.ID); err == nil {
		t.Fatalf("expected error retrieving deleted collection")
	}
}

func TestRouterCollectionsEndpoints(t *testing.T) {
	initDBForTest(t)
	defer db.DB.Close()
	// create collection via HTTP handler with valid body
	payload := []byte(`{"name":"router test","description":"test"}`)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/collections", bytes.NewBuffer(payload))
	CreateCollectionHandler(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

}
