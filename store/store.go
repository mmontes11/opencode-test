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

// DeleteItem removes an item by ID.
func DeleteItem(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM items WHERE id = ?", id)
	return err
}
