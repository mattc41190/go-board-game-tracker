package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Registers itself with the sql package
)

var USER = "root"
var PASSWORD = os.Args[1:][0]
var ADDRESS = "tcp(localhost:3306)/board_game_tracker_db"
var CONNECTION_STRING = fmt.Sprintf("%s:%s@%s", USER, PASSWORD, ADDRESS)

func getGameByTitle(title string) {

	db, err := sql.Open("mysql", CONNECTION_STRING)

	if err != nil {
		log.Fatal("Failed to connect", err)
	}

	defer db.Close()

	var s string
	query := "SELECT title FROM board_game_tracker_db.games WHERE title=?"
	err = db.QueryRow(query, title).Scan(&s)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("The query %s returned no rows", query)
		} else {
			log.Fatal("Failed to execute query", err)
		}
	}

	log.Println("Found:", s)

}
