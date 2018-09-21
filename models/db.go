package models

import (
	"database/sql"
	"fmt"
	"log"
)

// DataStore is an abstraction on top of the DB which allows simpler testing
type DataStore interface {
	GetAllGames() ([]*Game, error)
	SaveGame(Game)
}

// DB is a wrapper for sql.DB
type DB struct {
	*sql.DB
}

// NewDB gets a new DB for usage in application
func NewDB(user string, pass string, address string) (*DB, error) {

	connString := fmt.Sprintf("%s:%s@%s", user, pass, address)

	db, err := sql.Open("mysql", connString)

	if err != nil {
		log.Println("Failed to connect", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
