package models

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/mccaetano/cadSolidario/util"
)

type Scheduler struct {
	Id            int64           `json: id`
	EventDate     util.CustomTime `json: eventDate`
	EffectiveDate util.CustomTime `json: effectiveDate`
	Status        string          `json: status`
	Notes         string          `json: notes`
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

	row := DB.QueryRow("select id, event_date, effective_date, status, notes from public.tbcalendar where id = $1", id)
	var scheduler Scheduler
	row.Scan(&scheduler.Id, &scheduler.EventDate, &scheduler.EffectiveDate, &scheduler.Status, &scheduler.Notes)

	return scheduler, nil
}

func SchedulerGetByFilter(startEventDate time.Time, endEventDate time.Time, status string) ([]Scheduler, error) {

	rows, err := DB.Query(`select id, event_date, effective_date, status, notes 
		from public.tbcalendar where (event_date between $1 and $2 or '1900-01-01' = $1) and (status = $3 or '' = $3)`,
		startEventDate.Format("2006-01-02"),
		endEventDate.Format("2006-01-02"),
		status)
	if err != nil {
		log.Println("Erro lendo datados:", err)
		return nil, err
	}
	defer rows.Close()

	var schedulers []Scheduler
	for rows.Next() {
		var scheduler Scheduler
		rows.Scan(&scheduler.Id, &scheduler.EventDate, &scheduler.EffectiveDate, &scheduler.Status, &scheduler.Notes)
		schedulers = append(schedulers, scheduler)
	}

	return schedulers, nil
}

func SchedulerPost(scheduler Scheduler) (int64, error) {
	var id int64
	DB.QueryRow(`insert into public.tbcalendar ( event_date, effective_date, status, notes) 
		values ($1, $2, $3, $4)  RETURNING id`,
		scheduler.EventDate,
		scheduler.EffectiveDate,
		scheduler.Status,
		scheduler.Notes).Scan(&id)

	return id, nil
}

func SchedulerPut(id int64, scheduler Scheduler) (Scheduler, error) {

	_, err := DB.Exec(`update public.tbcalendar set event_date = $2, effective_date = $3, 
		status = $4, notes = $5 where id = $1`,
		id,
		scheduler.EventDate,
		scheduler.EffectiveDate,
		scheduler.Status,
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
