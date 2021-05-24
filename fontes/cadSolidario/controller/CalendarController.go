package controller

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mccaetano/cadSolidario/models"
)

func HandleGETCalendar(c *gin.Context) {
	cals, err := getByFilter(c.Query("startEventDate"), c.Query("endEventDate"), c.Query("status"))
	if err != nil {
		panic(err)
	}
	c.JSON(200, cals)
}

func HandlePOSTCalendar(c *gin.Context) {
	cals, err := getByFilter(c.Query("startEventDate"), c.Query("endEventDate"), c.Query("status"))
	if err != nil {
		panic(err)
	}
	c.JSON(200, cals)
}

func HandlePUTCalendar(c *gin.Context) {
	cals, err := getByFilter(c.Query("startEventDate"), c.Query("endEventDate"), c.Query("status"))
	if err != nil {
		panic(err)
	}
	c.JSON(200, cals)
}

/*-------- Internal -----------*/
func createDBConnecton() *sql.DB {
	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

func getByFilter(startEventDate string, endEventDate string, status string) ([]models.Calendar, error) {
	sqlStatament := "select * from tbCalendar "
	var sqlWhere string
	if startEventDate != "" && endEventDate != "" {
		if sqlWhere != "" {
			sqlWhere += " and "
		}
		sqlWhere += "eventDate between ? and ? "
	}
	if status != "" {
		if sqlWhere != "" {
			sqlWhere += " and "
		}
		sqlWhere += "status = ?"
	}
	if sqlWhere != "" {
		sqlStatament = sqlStatament + " where " + sqlWhere
	}

	db := createDBConnecton()

	defer db.Close()

	rows, err := db.Query(sqlStatament, startEventDate, endEventDate, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cals []models.Calendar
	for rows.Next() {
		var cal models.Calendar
		err := rows.Scan(&cal.Id, &cal.EventDate, &cal.Effective, &cal.Status, &cal.Notes)
		if err != nil {
			return nil, err
		}
		cals = append(cals, cal)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cals, nil

}
