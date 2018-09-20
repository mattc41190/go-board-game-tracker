package models

import (
	"database/sql"
)

// Game representation of a game
type Game struct {
	Name   string
	Played bool
	Link   string
	Rating int
}

// GetAllGames get all games from DB
func GetAllGames(db *sql.DB) ([]*Game, error) {
	rows, err := db.Query("SELECT title, played, link, rating FROM games")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	games := make([]*Game, 0)

	for rows.Next() {
		game := new(Game)
		err := rows.Scan(&game.Name, &game.Played, &game.Link, &game.Rating)

		if err != nil {
			return nil, err
		}

		games = append(games, game)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

// SaveGame saves a game to the database
func SaveGame(db *sql.DB, g Game) {
	// 1. Get user_id using the passed in password

	stmt := `INSERT INTO games (user_id, title, link, played, rating) VALUES (?, ?, ?, ?, ?);`
	_, err := db.Exec(stmt, 1, g.Name, g.Link, g.Played, g.Rating)
	if err != nil {
		panic(err)
	}
}
