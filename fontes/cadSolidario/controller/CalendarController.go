package api_controller

import (
	"database/sql"
    "github.com/mccaetano/cadSolidario/src/api/models"
	"github.com/gin-gonic/gin"
	"github.com/mccaetano/cadSolidario/models"
)

var DB *sql.DB

func handleGETCalendar(c *gin.Context) gin.ResponseWriter {
	cadCalendarREpository.DB := DB
	model := cadCalendarREpository.getByFilter(c.Query("startEventDate"), c.Query("endEventDate"), c.Query("status"))
	c.JSON(200,  model)
}
