package models

import (
	"database/sql"
	"fmt"
	"log"
)

// NewDB gets a new DB for usage in application
func NewDB(user string, pass string, address string) (*sql.DB, error) {

	connString := fmt.Sprintf("%s:%s@%s", user, pass, address)

	db, err := sql.Open("mysql", connString)

	if err != nil {
		log.Println("Failed to connect", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
