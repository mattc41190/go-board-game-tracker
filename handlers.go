package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattc41190/go-board-game-tracker/models"
)

func (e *Env) homePageHandler(c *gin.Context) {
	games, err := models.GetAllGames(e.db)

	if err != nil {
		log.Fatal(err)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"games": (games),
	})
}

func (e *Env) addGamePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "add-game.tmpl", gin.H{})
}

func (e *Env) addGameFormHandler(c *gin.Context) {
	secret := c.PostForm("secret")
	name := c.PostForm("name")
	link := c.PostForm("more-info-url")

	if secret != "PASS" {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"message": "You didn't put the right password in. Try again sometime :)",
		})
		return
	}

	g := models.Game{
		Name:   name,
		Link:   link,
		Played: false,
		Rating: 5,
	}

	fmt.Printf("Saving %v to the database", g.Name)
	models.SaveGame(e.db, g)

	c.Redirect(302, "/")

}
