package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Registers itself with the sql package
	"github.com/mattc41190/go-board-game-tracker/models"
)

var user = "root"
var pass = os.Args[1:][0]
var address = "tcp(localhost:3306)/board_game_tracker_db"
var connString = fmt.Sprintf("%s:%s@%s", user, pass, address)

// Env contains all global context
type Env struct {
	db models.DataStore
}

func main() {
	db, err := models.NewDB(user, pass, address)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	env := &Env{db}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", env.homePageHandler)
	r.GET("/add-game", env.addGamePageHandler)
	r.POST("/add-game", env.addGameFormHandler)

	r.Static("/assets", "./assets")

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
