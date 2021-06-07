package models

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Scheduler struct {
	Id        int64
	EventDate time.Time
	Effective time.Time
	Status    string
	Notes     string
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

	defer dbCon.Close()

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

	row := DB.QueryRow("select id, envent_date, sffective_date, status, notes from public.calendar where id = $1", id)

	var scheduler Scheduler
	row.Scan(&scheduler.Id, &scheduler.EventDate, &scheduler.Effective, &scheduler.Status, &scheduler.Notes)

	return scheduler, nil
}

func SchedulerGetByFilter(startEventDate time.Time, endEventDate time.Time, status string) ([]Scheduler, error) {

	rows, err := DB.Query(`select id, envent_date, sffective_date, status, notes 
		from public.calendar where (event_date between $1 and $2 or '0001-01-01' = $1) and (status = $3 ou '' = $3)`,
		startEventDate,
		endEventDate,
		status)
	if err != nil {
		log.Println("Erro lendo datados:", err)
		return nil, err
	}
	defer rows.Close()

	var schedulers []Scheduler
	for rows.Next() {
		var scheduler Scheduler
		rows.Scan(&scheduler.Id, &scheduler.EventDate, &scheduler.Effective, &scheduler.Status, &scheduler.Notes)
		schedulers = append(schedulers, scheduler)
	}

	return schedulers, nil
}

func SchedulerPost(scheduler Scheduler) (int64, error) {

	result, err := DB.Exec(`insert into public.calendar ( envent_date, sffective_date, status, notes) 
		values ($1, $2, $3, $4)`,
		scheduler.EventDate,
		scheduler.Effective,
		scheduler.Status,
		scheduler.Notes)
	if err != nil {
		log.Println("Erro lendo datados:", err)
		return 0, err
	}
	return result.LastInsertId()
}

func SchedulerPut(id int64, scheduler Scheduler) (Scheduler, error) {

	_, err := DB.Exec(`update public.calendar set envent_date = $2, effective_date = $3, 
		status = $4, notes = $5 where id = $1`,
		id,
		&scheduler.EventDate,
		&scheduler.Effective,
		&scheduler.Status,
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
