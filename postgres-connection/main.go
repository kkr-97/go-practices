package main

import (
	"database/sql"
	"log"
)

type PostgresServer struct {
	db *sql.DB
}

func (ps *PostgresServer) initDB() {
	db, err := sql.Open("postgres", "admin2:12345@localhost:5433/manage_rooms")
	if err != nil {
		log.Fatal(err)
	}
	ps.db = db
}

func main() {
	ps := &PostgresServer{}
	ps.initDB()
}
