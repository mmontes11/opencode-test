package db

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Init initializes the MariaDB connection. The DSN can be overridden by the
// MARIADB_DSN environment variable. A sample DSN: user:password@tcp(localhost:3306)/dbname
func Init() error {
    dsn := "root:password@tcp(localhost:3306)/mydb"
    // Allow override via env var for flexibility
    if envDSN := os.Getenv("MARIADB_DSN"); envDSN != "" {
        dsn = envDSN
    }
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("sql.Open: %w", err)
    }
    // Verify connection
    if err = DB.Ping(); err != nil {
        return fmt.Errorf("db ping: %w", err)
    }
    return nil
}
