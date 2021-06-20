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
	dateStart, err := time.Parse("2006-01-02", c.Query("startEventDate"))
	if err != nil {
		log.Printf("Erro to convert startEventDate(%s) - %s\n", c.Query("startEventDate"), err.Error())
		dateStart, _ = time.Parse("2006-01-02", "1900-01-01")
	}
	dateEnd, err := time.Parse("2006-01-02", c.Query("endEventDate"))
	if err != nil {
		log.Printf("Erro to convert endEventDate(%s) - %s\n", c.Query("endEventDate"), err.Error())
		dateEnd, _ = time.Parse("2006-01-02", "1900-01-01")
	}
	status := c.Query("status")
	limit, err := strconv.ParseInt(c.Query("limit"), 5, 32)
	if err != nil {
		log.Printf("Erro to convert limit(%s) - %s\n", c.Query("limit"), err.Error())
		limit = 20
	}

	skip, err := strconv.ParseInt(c.Query("skip"), 5, 32)
	if err != nil {
		log.Printf("Erro to convert skip(%s) - %s\n", c.Query("skip"), err.Error())
		skip = 0
	}
	log.Printf("Params: startEventDate: %s, endEventDate: %s, status: %s, Limit:%d, skip:%d\n", dateStart, dateEnd, status, limit, skip)
	cals, err := models.SchedulerGetByFilter(dateStart, dateEnd, status, int32(limit), int32(skip))
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
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Printf("Params: id: %d\n", id)
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
	err := c.BindJSON(&data)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Body In= %+v\n", data)
	id, err := models.SchedulerPost(data)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data.Id = id
	log.Printf("Finish: Body Out= %+v\n", data)
	c.Header("Content-type", "application/json")
	c.JSON(200, data)
}

func Put(c *gin.Context) {
	log.Println("Init: Put")
	var data models.Scheduler
	c.BindJSON(data)

	log.Printf("Finish: Body In= %+v\n", data)
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
