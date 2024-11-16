// internal/db/db.go
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Postgres driver
)

func NewDBConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open db connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping db: %v", err)
	}

	return db, nil
}
