package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/cadSolidario/models"
)

func GetStatusByFilter(c *gin.Context) {
	log.Println("Controller: (status) GetByFilter - Init")
	description := c.Query("description")
	limit, err := strconv.ParseInt(c.Query("limit"), 5, 32)
	if err != nil {
		log.Printf("Controller: (status) GetByFilter - Erro to convert limit(%s) - %s\n", c.Query("limit"), err.Error())
		limit = 20
	}

	skip, err := strconv.ParseInt(c.Query("skip"), 5, 32)
	if err != nil {
		log.Printf("Controller: (status) GetByFilter - Erro to convert skip(%s) - %s\n", c.Query("skip"), err.Error())
		skip = 0
	}
	log.Printf("Controller: (status) GetByFilter - Params= description:%s, Limit:%d, skip:%d\n", description, limit, skip)
	cals, err := models.StatusGetByFilter(description, int32(limit), int32(skip))
	if err != nil {
		log.Println("Controller: (status) GetByFilter - Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (status) GetByFilter - Finish: Body Out= %+v\n", cals)
	c.JSON(200, cals)
}

func GetStatuById(c *gin.Context) {
	log.Println("Controller: (status) GetById")

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Printf("Params: id: %d\n", id)
	cals, err := models.StatusGetById(id)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (status) GetByFilter - Finish: Body Out= %+v\n", cals)
	c.JSON(200, cals)
}

func PostPost(c *gin.Context) {
	log.Println("Controller: (status) - Post")
	var data models.Status
	err := c.BindJSON(&data)
	if err != nil {
		log.Println("Controller: (status) - Post: Error=", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Controller: (status) - Post: Body In= %+v\n", data)
	id, err := models.StatusPost(data)
	if err != nil {
		log.Println("Controller: (status) - Post: Error=", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data.Id = id
	log.Printf("Controller: (status) - Post: Body Out= %+v\n", data)
	c.Header("Content-type", "application/json")
	c.JSON(200, data)
}

func PutStatus(c *gin.Context) {
	log.Println("Controller: (status) - Put: Init")
	var data models.Status
	c.BindJSON(data)

	log.Printf("Controller: (scheduler) - Put: Body In= %+v\n", data)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cal, err := models.StatusPut(id, data)
	if err != nil {
		log.Println("Controller: (status) - Put: Error=", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (status) - Put: Body Out= %+v\n", cal)
	c.JSON(200, cal)
}
