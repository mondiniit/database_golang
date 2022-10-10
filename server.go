package database_golang

import (
	game_application "database_golang/src/application"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/", game_application.ResponseGame)

	router.Static("/assets", "./assets")
	router.StaticFile("/styles.css", "./assets/styles.css")
	router.StaticFile("/star.png", "./assets/star.png")
	router.StaticFile("/main.js", "./assets/main.js")
	router.LoadHTMLGlob("assets/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Run("localhost:8080")
}
