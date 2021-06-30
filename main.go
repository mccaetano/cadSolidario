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
	api.GET("/calendar", controller.GetCalendarByFilter)
	api.GET("/calendar/event-dates", controller.GetCalendarEventDate)
	api.POST("/calendar", controller.PostCalendar)
	api.PUT("/calendar/:id", controller.PutCalendar)

	api.GET("/status", controller.GetStatusByFilter)
	api.GET("/status/:id", controller.GetStatuById)
	api.POST("/status", controller.PostPost)
	api.PUT("/status", controller.PutStatus)

	api.GET("/recipient", controller.GetRecipientByFilter)
	api.GET("/recipient/:id", controller.GetRecipientById)
	api.POST("/recipient", controller.PostCalendar)
	api.PUT("/recipient/:id", controller.PutCalendar)

	api.GET("/address/:postalCode", controller.GetStatusByFilter)
	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	router.Run(":" + port)
}
