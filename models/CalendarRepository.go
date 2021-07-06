package models

import (
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Calendar struct {
	Id        int      `json:"id,omitempty"`
	EventDate string   `json:"eventDate,omitempty"`
	Name      string   `json:"name,omitempty"`
	Documents Document `json:"documents,omitempty"`
	Contacts  Contact  `json:"contacts,omitempty"`
	Milks     int      `json:"milks,omitempty"`
	Status    Status   `json:"status,omitempty"`
}

func CalendarGetByFilter(eventDate time.Time, name string, limit int32, skip int32) ([]Calendar, error) {
	offset := limit * (skip - 1)
	log.Println("Repo: (Calendar) Get - eventDate: ", eventDate.Format("2006-01-02"), ", name: ", name, ", limit: ", limit, ", skip: ", skip)
	rows, err := DB.Query(
		`select 
				t.id ,
				b.nome,
				b.documento_cpf,
				b.contato_tel,
				b.contato_cel,
				b.qtde_leite,
				s.id status_id,
				s.descricao status_descricao	
			from 	
				tbagendamento t 
				inner join tbbeneficiario b on
					b.id  = t.beneficiario 
				inner  join tbstatus s on 
					s.id  = t.status
			where 
				t.event_date = $1
				and b.nome like '%' || $2 || '%'
			order by 
				t.event_date , b.nome 
			limit $3 offset  $4`,
		eventDate.Format("2006-01-02"),
		name,
		limit,
		offset)
	if err != nil {
		log.Println("Repo: (Calendar) Get - Erro: ", err)
		return nil, err
	}
	defer rows.Close()

	var calendars []Calendar = []Calendar{}
	for rows.Next() {
		var calendar Calendar
		rows.Scan(&calendar.Id, &calendar.Name, &calendar.Documents.Cpf, &calendar.Contacts.Phone,
			&calendar.Contacts.CelPhone, &calendar.Milks, &calendar.Status.Id, &calendar.Status.Description)
		calendars = append(calendars, calendar)
	}
	log.Printf("Repo: (Calendar) Get - body Out= %+v\n", calendars)
	return calendars, nil
}

func CalendarGetEventDates() ([]Calendar, error) {
	log.Println("Repo: (Calendar) Get - CalendarGetEventDates: ")
	rows, err := DB.Query(
		`select distinct 
			to_char(t.event_date, 'YYYY-MM-DD') event_date
		from
			tbagendamento t 
		order by
			event_date desc`)
	if err != nil {
		log.Println("Repo: (Calendar) Get - Erro: ", err)
		return nil, err
	}
	defer rows.Close()

	var calendars []Calendar = []Calendar{}
	for rows.Next() {
		var calendar Calendar
		rows.Scan(&calendar.EventDate)
		calendars = append(calendars, calendar)
	}
	log.Printf("Repo: (Calendar) Get - body Out= %+v\n", calendars)
	return calendars, nil
}

func CalendarPost(eventDate time.Time) error {
	log.Printf("Repo: (Calendar) Post - eventDate: %s\n", eventDate.Format("2006-01-02"))
	stmt, err := DB.Prepare(`insert into tbagendamento (event_date, status, beneficiario)
			select
				$1 as "event_date",
				case
					when a.status = 8 then 8
					else 7
				end status,
				b.id beneficiario
			from
				tbbeneficiario b
			left join tbagendamento a on
				a.beneficiario = b.id
				and a.event_date = (select max(c.event_date) from tbagendamento c)
			where
				b.ativo = true`)
	if err != nil {
		log.Println("Repo: (Calendar) Post - Erro: ", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(eventDate.Format("2006-01-02"))
	if err != nil {
		log.Println("Repo: (Calendar) Post - Erro: ", err)
		return err
	}
	log.Printf("Repo: (Calendar) Post - count= %+v\n", result)
	return nil
}

func CalendarPut(id int, status int) error {
	log.Printf("Repo: (Calendar) Put - id: %d, status: %d\n", id, status)
	_, err := DB.Exec(`update tbagendamento set status = $2 where id = $1`,
		id,
		status)
	if err != nil {
		log.Println("Repo: (Calendar) Put - Erro lendo datados:", err)
		return err
	}
	log.Printf("Repo: (Calendar) Put - \n")
	return nil
}
