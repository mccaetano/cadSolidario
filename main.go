package main

import (
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/mccaetano/cadSolidario/controller"
	"github.com/mccaetano/cadSolidario/models"
)

func main() {

	err := models.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())

	api := router.Group("/api")
	api.GET("/calendar/status", controller.GetByStatus)
	api.GET("/calendar", controller.GetByFilter)
	api.GET("/calendar/:id", controller.GetById)
	api.POST("/calendar", controller.Post)
	api.PUT("/calendar/:id", controller.Put)

	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	router.Run(":" + port)
}
