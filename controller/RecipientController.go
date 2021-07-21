package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/cadSolidario/models"
)

func GetRecipientByFilter(c *gin.Context) {
	log.Println("Controller: (Recipient) GetAddressByPostalCode")

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

	log.Printf("Params: name: %s\n", name)
	cals, err := models.RecipientGetByFilter(name, int32(limit), int32(skip))
	if err != nil {
		log.Println("Controller: (Recipient) - Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (Recipient) GetAddressByPostalCode - Finish: Body Out= %+v\n", cals)
	c.JSON(200, cals)
}

func GetRecipientById(c *gin.Context) {
	log.Println("Controller: (Recipient) GetRecipientById")

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Printf("Params: id: %d\n", id)
	cals, err := models.RecipientGetById(id)
	if err != nil {
		log.Println("Controller: (Recipient) - Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (Recipient) GetRecipientById - Finish: Body Out= %+v\n", cals)
	c.JSON(200, cals)
}

func PostRecipient(c *gin.Context) {
	log.Println("Controller: (recipient) - Post")
	var data models.Recipient
	err := c.BindJSON(&data)
	if err != nil {
		log.Println("Controller: (recipient) - Post: Error=", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (recipient) - Post: Body In= %+v\n", data)

	id, err := models.RecipientPost(data)
	if err != nil {
		log.Println("Controller: (recipient) - Post: Error=", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Controller: (recipient) - Post: Body Out= %d\n", id)
	c.Header("Content-type", "application/json")
	c.JSON(201, gin.H{})
}

func PutRecipient(c *gin.Context) {
	log.Println("Controller: (recipient) - Put: Init")
	var data models.Recipient
	c.BindJSON(&data)

	log.Printf("Controller: (recipient) - Put: Body In= %+v\n", data)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	recipient, err := models.RecipientPut(id, data)
	if err != nil {
		log.Println("Controller: (recipient) - Put: Error=", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Controller: (recipient) - Put: Body Out= \n")

	c.JSON(200, recipient)
}
