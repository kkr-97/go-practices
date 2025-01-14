package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var db *sql.DB

func init() {
	var err error
	dsn := "postgres://postgres:durga%401234@localhost:5432/manage_rooms?sslmode=disable"
	db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

}

func generateCode() string {
	code := make([]byte, 4)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

func createRoom(c *gin.Context) {
	var code string
	err := db.QueryRow("SELECT code FROM rooms WHERE status = 'available' LIMIT 1").Scan(&code)

	if err == sql.ErrNoRows {
		code = generateCode() //edge case 1: duplicate code generated
		_, err = db.Exec("INSERT INTO rooms (code, status) VALUES ($1, $2)", code, "unavailable")
		if err != nil {
			fmt.Println("Error while inserting new code:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
			return
		}
	} else if err != nil {
		fmt.Println("DatabaseError:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	} else {
		_, err = db.Exec("UPDATE rooms SET status = 'unavailable' WHERE code = $1", code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room status"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"room_code": code})
}

func releaseRoom(c *gin.Context) {
	code := c.Param("code")

	// Check if the room exists and is not already released
	var status string
	err := db.QueryRow("SELECT status FROM rooms WHERE code = $1", code).Scan(&status)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found to release"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query room status"})
		return
	}

	if status == "available" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room is already released"})
		return
	}

	// Update the room status to 'available'
	_, err = db.Exec("UPDATE rooms SET status = 'available' WHERE code = $1", code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to release room"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room released successfully"})
}

func main() {
	r := gin.Default()

	r.POST("/createroom", createRoom)
	r.PUT("/releaseroom/:code", releaseRoom)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
