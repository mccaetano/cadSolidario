package models

import (
	"database/sql"

	"log"
	"os"

	_ "github.com/lib/pq"
)

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
