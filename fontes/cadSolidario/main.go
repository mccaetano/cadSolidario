package main

import (
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/mccaetano/cadSolidario/controller"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())

	api := router.Group("/api")
	api.GET("/calendar", controller.HandleGETCalendar)
	api.POST("/calendar", controller.HandlePOSTCalendar)
	api.PUT("/calendar/{id}", controller.HandlePUTCalendar)

	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	router.Run(":" + port)
}
