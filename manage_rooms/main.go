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
		code = generateCode() //edge case: duplicate code generated
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

	_, err := db.Exec("UPDATE rooms SET status = 'available' WHERE code = $1", code)
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
