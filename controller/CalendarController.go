package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/cadSolidario/models"
)

func GetByFilter(c *gin.Context) {
	log.Println("Init: GetByFilter")
	dateStart, err := time.Parse("2006-03-01", c.Query("startEventDate"))
	if err != nil {
		dateStart, _ = time.Parse("2006-03-01", "1900-01-01")
	}
	dateEnd, err := time.Parse("2006-03-01", c.Query("endEventDate"))
	if err != nil {
		dateEnd, _ = time.Parse("2006-03-01", "1900-01-01")
	}
	log.Println("Params: Post")
	cals, err := models.SchedulerGetByFilter(dateStart, dateEnd, c.Query("status"))
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Finish: Body Out= %+v\n", cals)
	c.JSON(200, cals)
}

func GetById(c *gin.Context) {
	log.Println("Init: GetById")
	id, _ := strconv.ParseInt(c.Query("status"), 10, 64)
	cals, err := models.SchedulerGetById(id)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Finish: Body Out= %+v\n", cals)
	c.JSON(200, cals)
}

func Post(c *gin.Context) {
	log.Println("Init: Post")
	var data models.Scheduler
	c.BindJSON(&data)
	cals, err := models.SchedulerPost(data)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Finish: Body Out= %+v\n", cals)
	c.JSON(200, cals)
}

func Put(c *gin.Context) {
	log.Println("Init: Put")
	var data models.Scheduler
	c.BindJSON(&data)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cal, err := models.SchedulerPut(id, data)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Finish: Body Out= %+v\n", cal)
	c.JSON(200, cal)
}
