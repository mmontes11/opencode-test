package store

import (
	"database/sql"
	// Removed time import
)

// Item represents a simple record in the items table.
// The table schema is:
//   id BIGINT PRIMARY KEY AUTO_INCREMENT,
//   name VARCHAR(255) NOT NULL,
//   description TEXT,
//   created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
// Note: The actual migration is handled externally; this file simply provides an API.

// Item holds item data returned to clients.
// The CreatedAt field is kept as a string to avoid timezone parsing complexities.
// Tests assert the string value directly.
type Item struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   string
}

// CreateItem inserts a new item into the database and returns its details.
func CreateItem(db *sql.DB, name, description string) (*Item, error) {
	res, err := db.Exec("INSERT INTO items (name, description) VALUES (?, ?)", name, description)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT id, name, description, created_at FROM items WHERE id = ?", id)
	var item Item
	if err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt); err != nil {
		return nil, err
	}
	return &item, nil
}

// GetItem retrieves an item by its ID.
func GetItem(db *sql.DB, id int64) (*Item, error) {
	row := db.QueryRow("SELECT id, name, description, created_at FROM items WHERE id = ?", id)
	var item Item
	if err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt); err != nil {
		return nil, err
	}
	return &item, nil
}

// ListItems returns all items.
func ListItems(db *sql.DB) ([]Item, error) {
	rows, err := db.Query("SELECT id, name, description, created_at FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var itm Item
		if err := rows.Scan(&itm.ID, &itm.Name, &itm.Description, &itm.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, itm)
	}
	return items, rows.Err()
}

// UpdateItem modifies an existing item.
func UpdateItem(db *sql.DB, id int64, name, description string) (*Item, error) {
	_, err := db.Exec("UPDATE items SET name = ?, description = ? WHERE id = ?", name, description, id)
	if err != nil {
		return nil, err
	}
	return GetItem(db, id)
}

func DeleteItem(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM items WHERE id = ?", id)
	return err
}

// Collection represents a collection of items.
// The table schema is:
//   id BIGINT PRIMARY KEY AUTO_INCREMENT,
//   name VARCHAR(255) NOT NULL,
//   description TEXT,
//   created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
// Note: The actual migration is handled externally; this file simply provides an API.

// Collection holds collection data returned to clients.
// The CreatedAt field is kept as a string to avoid timezone parsing complexities.
// Tests would assert the string value directly.
type Collection struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   string
}

// CreateCollection inserts a new collection into the database and returns its details.
func CreateCollection(db *sql.DB, name, description string) (*Collection, error) {
	res, err := db.Exec("INSERT INTO collections (name, description) VALUES (?, ?)", name, description)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT id, name, description, created_at FROM collections WHERE id = ?", id)
	var col Collection
	if err := row.Scan(&col.ID, &col.Name, &col.Description, &col.CreatedAt); err != nil {
		return nil, err
	}
	return &col, nil
}

// GetCollection retrieves a collection by its ID.
func GetCollection(db *sql.DB, id int64) (*Collection, error) {
	row := db.QueryRow("SELECT id, name, description, created_at FROM collections WHERE id = ?", id)
	var col Collection
	if err := row.Scan(&col.ID, &col.Name, &col.Description, &col.CreatedAt); err != nil {
		return nil, err
	}
	return &col, nil
}

// ListCollections returns all collections.
func ListCollections(db *sql.DB) ([]Collection, error) {
	rows, err := db.Query("SELECT id, name, description, created_at FROM collections")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cols []Collection
	for rows.Next() {
		var col Collection
		if err := rows.Scan(&col.ID, &col.Name, &col.Description, &col.CreatedAt); err != nil {
			return nil, err
		}
		cols = append(cols, col)
	}
	return cols, rows.Err()
}

// UpdateCollection modifies an existing collection.
func UpdateCollection(db *sql.DB, id int64, name, description string) (*Collection, error) {
	_, err := db.Exec("UPDATE collections SET name = ?, description = ? WHERE id = ?", name, description, id)
	if err != nil {
		return nil, err
	}
	return GetCollection(db, id)
}

// DeleteCollection removes a collection by ID, and cleans up relationships.
func DeleteCollection(db *sql.DB, id int64) error {
	// Remove from join table first
	if _, err := db.Exec("DELETE FROM collection_items WHERE collection_id = ?", id); err != nil {
		return err
	}
	_, err := db.Exec("DELETE FROM collections WHERE id = ?", id)
	return err
}

// AddItemToCollection associates an item with a collection.
func AddItemToCollection(db *sql.DB, collectionID, itemID int64) error {
	_, err := db.Exec("INSERT IGNORE INTO collection_items (collection_id, item_id) VALUES (?, ?)", collectionID, itemID)
	return err
}

// ListItemsInCollection retrieves all items belonging to the specified collection.
func ListItemsInCollection(db *sql.DB, collectionID int64) ([]Item, error) {
	rows, err := db.Query("SELECT i.id, i.name, i.description, i.created_at FROM items i JOIN collection_items ci ON i.id = ci.item_id WHERE ci.collection_id = ?", collectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var itm Item
		if err := rows.Scan(&itm.ID, &itm.Name, &itm.Description, &itm.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, itm)
	}
	return items, rows.Err()
}

// RemoveItemFromCollection disassociates an item from a collection.
func RemoveItemFromCollection(db *sql.DB, collectionID, itemID int64) error {
	_, err := db.Exec("DELETE FROM collection_items WHERE collection_id = ? AND item_id = ?", collectionID, itemID)
	return err
}

// Duplicate DeleteItem removed
