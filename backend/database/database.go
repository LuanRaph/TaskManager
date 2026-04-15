package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := "host=localhost user=postgres password=YOUR_PASSWORD dbname=postgres sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Banco não respondeu", err)
	}
	log.Println("Banco conectado")
}
