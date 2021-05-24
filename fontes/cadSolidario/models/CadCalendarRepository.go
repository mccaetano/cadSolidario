package api_model

import (
	"database/sql"
	"time"
)

var DB *sql.DB

type calendar struct {
	id        int64
	eventDate time.Time
	effective time.Time
	status    string
	notes     string
}

func getByFilter(startEventDate time.Time, endEventDate time.Time, status string) ([]calendar, error) {
	sqlStatament := "select * from tbCalendar "
	var sqlWhere string
	if !startEventDate.IsZero() && !endEventDate.IsZero() {
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

	rows, err := DB.Query(sqlStatament, startEventDate.Format("yyyy-MM-DD"), endEventDate.Format("yyyy-MM-dd"), status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cals []calendar
	for rows.Next() {
		var cal calendar
		err := rows.Scan(&cal.id, &cal.eventDate, &cal.effective, &cal.status, &cal.notes)
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
