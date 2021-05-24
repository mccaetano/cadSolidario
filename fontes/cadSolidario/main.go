package main

import (
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.Group("/api")

	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	router.Run(":" + port)
}
