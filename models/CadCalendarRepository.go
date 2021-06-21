package models

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Status struct {
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
}

type Scheduler struct {
	Id            int64  `json:"id"`
	EventDate     string `json:"eventDate,omitempty"`
	EffectiveDate string `json:"effectiveDate,omitempty"`
	Sta           Status `json:"status,omitempty"`
	Notes         string `json:"notes,omitempty"`
}

var DB *sql.DB

func ConnectDatabase() error {
	url := os.Getenv("DATABASE_URL")

	log.Println("Connecting: ", url)

	dbCon, err := sql.Open("postgres", url)
	if err != nil {
		log.Println("Erro ao connectar no banco de dados: ", url)
		return err
	}

	//defer dbCon.Close()

	// check the connection
	err = dbCon.Ping()

	if err != nil {
		log.Println("Erro ao connectar no banco de dados: ", url)
		return err
	}

	DB = dbCon

	return nil
}

func SchedulerGetById(id int64) (Scheduler, error) {

	row := DB.QueryRow(`select
				c.id,
				to_char(c.event_date, 'YYYY-MM-DD') as event_date,
				to_char(c.effective_date, 'YYYY-MM-DD') as effective_date,
				c.status,
				s.description as status_description,
				c.notes
			from
				public.tbcalendar c
				inner join public.tbstatus s on
					s.code = c.status 
			where
				c.id = $1`, id)
	var scheduler Scheduler
	row.Scan(&scheduler.Id, &scheduler.EventDate, &scheduler.EffectiveDate, &scheduler.Sta.Code,
		&scheduler.Sta.Description, &scheduler.Notes)

	return scheduler, nil
}

func SchedulerGetByFilter(startEventDate time.Time, endEventDate time.Time, status string, limit int32, skip int32) ([]Scheduler, error) {
	offset := limit * (skip - 1)
	rows, err := DB.Query(`select
				c.id,
				to_char(c.event_date, 'YYYY-MM-DD') as event_date,
				to_char(c.effective_date, 'YYYY-MM-DD') as effective_date,
				c.status,
				s.description as status_description,
				c.notes
			from
				public.tbcalendar c
				inner join public.tbstatus s on
					s.code = c.status
			where
				(c.event_date between $1 and $2
					or '1900-01-01' = $1)
				and (c.status = $3
					or '' = $3)
			order by
				c.event_date desc,
				c.id desc
			limit $4 offset $5`,
		startEventDate.Format("2006-01-02"),
		endEventDate.Format("2006-01-02"),
		status,
		limit,
		offset)
	if err != nil {
		log.Println("Erro lendo datados:", err)
		return nil, err
	}
	defer rows.Close()

	var schedulers []Scheduler = []Scheduler{}
	for rows.Next() {
		var scheduler Scheduler
		rows.Scan(&scheduler.Id, &scheduler.EventDate, &scheduler.EffectiveDate, &scheduler.Sta.Code,
			&scheduler.Sta.Description, &scheduler.Notes)
		if string(scheduler.EffectiveDate) == "0001-01-01" || scheduler.EffectiveDate == "1900-01-01" {
			scheduler.EffectiveDate = ""
		}
		if string(scheduler.EventDate) == "0001-01-01" || scheduler.EventDate == "1900-01-01" {
			scheduler.EventDate = ""
		}
		schedulers = append(schedulers, scheduler)
	}

	return schedulers, nil
}

func SchedulerGetStatus() ([]Status, error) {
	rows, err := DB.Query(`select
			s.code ,
			s.description 
		from
			public.tbstatus s`)
	if err != nil {
		log.Println("Erro lendo datados:", err)
		return nil, err
	}
	defer rows.Close()

	var sstatuss []Status = []Status{}
	for rows.Next() {
		var status Status
		rows.Scan(&status.Code, &status.Description)
		sstatuss = append(sstatuss, status)
	}

	return sstatuss, nil
}

func SchedulerPost(scheduler Scheduler) (int64, error) {
	var id int64
	var eventDate string = "1900-01-01"
	if scheduler.EventDate != "" {
		date, err := time.Parse("2006-01-02", scheduler.EventDate)
		if err != nil {
			eventDate = date.Format("2006-01-02")
		}
	}
	var effectiveDate string = "1900-01-01"
	if scheduler.EffectiveDate != "" {
		date, err := time.Parse("2006-01-02", scheduler.EffectiveDate)
		if err != nil {
			effectiveDate = date.Format("2006-01-02")
		}
	}

	DB.QueryRow(`insert into public.tbcalendar ( event_date, effective_date, status, notes) 
		values ($1, $2, $3, $4)  RETURNING id`,
		eventDate,
		effectiveDate,
		scheduler.Sta.Code,
		scheduler.Notes).Scan(&id)

	return id, nil
}

func SchedulerPut(id int64, scheduler Scheduler) (Scheduler, error) {

	_, err := DB.Exec(`update public.tbcalendar set event_date = $2, effective_date = $3, 
		status = $4, notes = $5 where id = $1`,
		id,
		scheduler.EventDate,
		scheduler.EffectiveDate,
		scheduler.Sta.Code,
		&scheduler.Notes)
	if err != nil {
		log.Println("Erro lendo datados:", err)
		return Scheduler{}, err
	}

	schedulerAlt, err := SchedulerGetById(id)
	if err != nil {
		log.Println("Erro lendo datados:", err)
		return Scheduler{}, err
	}
	return schedulerAlt, err
}
