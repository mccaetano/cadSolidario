package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/cadSolidario/models"
)

type PostCalendarRequest struct {
	EventDate string `json:"eventDate"`
}
type PutCalendarRequest struct {
	Status models.Status `json:"status"`
}

func GetCalendarByFilter(c *gin.Context) {
	eventDate, err := time.Parse("2006-01-02", c.Query("eventDate"))
	if err != nil {
		log.Printf("Controller: (calendar) GetByFilter - Erro to convert eventDate(%s) - %s\n", c.Query("eventDate"), err.Error())
		eventDate, _ = time.Parse("2006-01-02", "1900-01-01")
	}
	name := c.Query("name")
	limit, err := strconv.ParseInt(c.Query("limit"), 5, 32)
	if err != nil {
		log.Printf("Controller: (calendar) GetByFilter - Erro to convert limit(%s) - %s\n", c.Query("limit"), err.Error())
		limit = 20
	}
	skip, err := strconv.ParseInt(c.Query("skip"), 5, 32)
	if err != nil {
		log.Printf("Controller: (calendar) GetByFilter - Erro to convert skip(%s) - %s\n", c.Query("skip"), err.Error())
		skip = 0
	}
	log.Printf("Controller: (calendar) GetByFilter - Params: eventDate: %s, name: %s, Limit:%d, skip:%d\n", eventDate, name, limit, skip)
	cals, err := models.CalendarGetByFilter(eventDate, name, int32(limit), int32(skip))
	if err != nil {
		log.Println("Controller: (calendar) GetByFilter - Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (calendar) GetByFilter - Finish: Body Out= %+v\n", cals)
	c.JSON(200, cals)
}

func GetCalendarEventDate(c *gin.Context) {
	log.Printf("Controller: (calendar) GetCalendarEventDate - \n")
	cals, err := models.CalendarGetEventDates()
	if err != nil {
		log.Println("Controller: (calendar) GetCalendarEventDate - Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (calendar) GetCalendarEventDate - Finish: Body Out= %+v\n", cals)
	c.JSON(200, cals)
}

func PostCalendar(c *gin.Context) {
	log.Println("Controller: (calendar) - Post")
	var data PostCalendarRequest
	err := c.BindJSON(&data)
	if err != nil {
		log.Println("Controller: (calendar) - Post: Error=", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (calendar) - Post: Body In= %+v\n", data)

	eventDate, err := time.Parse("2006-01-02", data.EventDate)
	if err != nil {
		log.Printf("Controller: (calendar) Post - Erro to convert eventDate(%s) - %s\n", data.EventDate, err.Error())
		eventDate, _ = time.Parse("2006-01-02", "1900-01-01")
	}
	err = models.CalendarPost(eventDate)
	if err != nil {
		log.Println("Controller: (calendar) - Post: Error=", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Controller: (calendar) - Post: Body Out= \n")
	c.Header("Content-type", "application/json")
	c.JSON(201, gin.H{})
}

func PutCalendar(c *gin.Context) {
	log.Println("Controller: (calendar) - Put: Init")
	var data PutCalendarRequest
	c.BindJSON(data)

	log.Printf("Controller: (calendar) - Post: Body In= %+v\n", data)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := models.CalendarPut(id, data.Status.Id)
	if err != nil {
		log.Println("Controller: (calendar) - Post: Error=", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (calendar) - Post: Body Out= \n")
	c.JSON(201, gin.H{})
}
