package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func initDB() {
	var err error
	dsn := "postgres://admin2:12345@localhost:5433/manage_rooms?sslmode=disable"
	db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS rooms (
            code INTEGER PRIMARY KEY,
            is_available BOOLEAN
        );
    `)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Check if the rooms table is empty
	var rowCount int
	err = db.QueryRow("SELECT COUNT(*) FROM rooms").Scan(&rowCount)
	if err != nil {
		log.Fatalf("Failed to check row count: %v", err)
	}

	// If the table is empty, insert rows
	if rowCount == 0 {
		_, err = db.Exec(`
			INSERT INTO rooms (code, is_available)
			SELECT generate_series(0, 1679615), true;
		`)
		if err != nil {
			log.Fatalf("Failed to insert room data: %v", err)
		}
	}
}

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base36ToInt(code string) int {
	code = strings.ToUpper(code) // Ensure it's uppercase
	value := 0
	for i, char := range code {
		pos := strings.IndexRune(charset, char) // Find position in charset
		if pos == -1 {
			panic(fmt.Sprintf("Invalid character: %c", char))
		}
		exp := 3 - i
		value += pos * int(math.Pow(36, float64(exp)))
	}
	return value
}

// IntToBase36 converts an integer to a Base-36 string
func IntToBase36(num int) string {
	if num < 0 {
		panic("Negative numbers are not allowed")
	}
	result := ""
	for num > 0 {
		ind := num % 36
		result = string(charset[ind]) + result
		num /= 36
	}
	for len(result) < 4 {
		result = "0" + result
	}
	return result
}

func createRoom(c *gin.Context) {
	var int_code int
	err := db.QueryRow("SELECT code FROM rooms WHERE is_available = true LIMIT 1").Scan(&int_code)

	code := IntToBase36(int_code)
	if err == sql.ErrNoRows {
		log.Fatal("No available codes, please wait")
	} else if err != nil {
		fmt.Println("DatabaseError:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	} else {
		_, err = db.Exec("UPDATE rooms SET is_available = false WHERE code = $1", int_code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room is_available", "desc": err})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"room_code": code})
}

func releaseRoom(c *gin.Context) {
	code := Base36ToInt(c.Param("code"))

	// Check if the room exists and is not already released
	var is_available bool
	err := db.QueryRow("SELECT is_available FROM rooms WHERE code = $1", code).Scan(&is_available)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found to release"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query room is_available"})
		return
	}

	fmt.Println(is_available, code)

	if is_available {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room is already released"})
		return
	}

	// Update the room is_available to 'available'
	_, err = db.Exec("UPDATE rooms SET is_available = true WHERE code = $1", code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to release room"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room released successfully"})
}

func main() {
	initDB()
	r := gin.Default()

	r.PUT("/createroom", createRoom)
	r.PUT("/releaseroom/:code", releaseRoom)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
