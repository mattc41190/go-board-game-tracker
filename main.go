package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	App := gin.Default()
	App.LoadHTMLGlob("templates/*")

	getGameByTitle("Nefarious")

	// addRoutes(App)
	// App.Static("/assets", "./assets")

	// App.Run(":8080") // listen and serve on 0.0.0.0:8080
}
