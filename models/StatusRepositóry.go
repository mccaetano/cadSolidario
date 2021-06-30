package models

import (
	"log"
)

type Status struct {
	Id          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

func StatusGetById(id int64) (Status, error) {
	log.Println("Repo: (Status) Get - id: ", id)

	row := DB.QueryRow(`select 
			t.id,
			t.descricao 
		from
			tbstatus t 
		where 
			t.id  = $1`, id)
	var status Status
	row.Scan(&status.Id, &status.Description)
	log.Printf("Repo: (Status) Get - body Out= %+v\n", status)
	return status, nil
}

func StatusGetByFilter(description string, limit int32, skip int32) ([]Status, error) {
	offset := limit * (skip - 1)
	log.Println("Repo: (Status) Get - description: ", description, ", limit: ", limit, ", skip: ", skip)
	rows, err := DB.Query(`select 
				t.id,
				t.descricao 
			from
				tbstatus t 
			where 
				t.descricao like '%' || $1 || '%'
			limit $2 offset $3`,
		description,
		limit,
		offset)
	if err != nil {
		log.Println("Repo: (Status) Get - Erro lendo datados:", err)
		return nil, err
	}
	defer rows.Close()

	var statuss []Status = []Status{}
	for rows.Next() {
		var status Status
		rows.Scan(&status.Id, &status.Description)
		statuss = append(statuss, status)
	}
	log.Printf("Repo: (Status) Post - body Out= %+v\n", statuss)
	return statuss, nil
}

func StatusPost(status Status) (int, error) {
	log.Printf("Repo: (Status) Post - body: %+v\n", status)

	var id int
	DB.QueryRow(`INSERT INTO public.tbstatus
			(descricao)
			VALUES($1)
				RETURNING id`,
		status.Description).Scan(&id)
	log.Println("Repo: (Status) Post - id: ", id)
	return id, nil
}

func StatusPut(id int64, status Status) (Status, error) {
	log.Printf("Repo: (Status) Put - id: %d, body: %+v\n", id, status)
	_, err := DB.Exec(`UPDATE public.tbstatus
			SET descricao=$2
			WHERE id=$1;
			`,
		id,
		status.Description)
	if err != nil {
		log.Println("Repo: (Status) Put - Erro lendo datados:", err)
		return Status{}, err
	}

	statusAlt, err := StatusGetById(id)
	if err != nil {
		log.Println("Repo: (Status) Put - Erro lendo datados:", err)
		return Status{}, err
	}
	log.Printf("Repo: (Status) Put - body: %+v\n", statusAlt)
	return statusAlt, err
}
