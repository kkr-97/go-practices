package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresServer struct {
	db *sql.DB
}

func (ps *PostgresServer) ping() error {
	return ps.db.Ping()
}

func (ps *PostgresServer) initDB() {
	// dsn := "postgres://admin2:12345@localhost:5433/manage_rooms" //local
	dsn := "postgres://admin2:12345@postgres-db:5432/manage_rooms?" //docker
	db, _ := sql.Open("postgres", dsn)
	ps.db = db
	if err := ps.ping(); err != nil {
		log.Fatal("DB not connected...", err)
	}
	fmt.Println("Database connected successfully!")
	go ps.monitorDBConnection()
}

func (ps *PostgresServer) monitorDBConnection() {
	for {

		if err := ps.ping(); err != nil {
			fmt.Println("DB connection interrupted, trying to reconnect...:", err)
			ok := ps.reconnect()
			if !ok {
				log.Fatal("Failed to reconnect to DB")
			}
		} else {
			fmt.Println("DB connection is active")
		}

		time.Sleep(5 * time.Second)
	}
}

func (ps *PostgresServer) reconnect() bool {
	fmt.Println("Attempting to reconnect to PostgreSQL...")
	cwt, _ := context.WithTimeout(context.Background(), 20*time.Second)
	for {
		select {
		case <-cwt.Done():
			return false
		default:

			if err := ps.ping(); err == nil {
				fmt.Println("DB re-connected!!")
				return true
			} else {
				fmt.Println("Still Retrying...")
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ps := &PostgresServer{}
	ps.initDB()

	select {}
}
