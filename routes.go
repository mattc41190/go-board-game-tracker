package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addRoutes(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"games": getGamesFromDB(),
		})
	})

	app.GET("/add-game", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add-game.tmpl", gin.H{})
	})

	app.POST("/add-game", func(c *gin.Context) {
		secret := c.PostForm("secret")
		name := c.PostForm("name")
		link := c.PostForm("more-info-url")

		if secret != "PASS" {
			c.HTML(http.StatusOK, "error.tmpl", gin.H{
				"message": "You didn't put the right password in. Try again sometime :)",
			})
			return
		}

		Game := game{
			Name:   name,
			Link:   link,
			Played: false,
			Rating: 5,
		}

		fmt.Printf("Saving %v to the database", Game.Name)

		c.Request.URL.Path = "/"
		c.Request.Method = "GET"
		app.HandleContext(c)

	})
}
