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
