package controllers

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/kkr-97/manage_rooms/services"
// )

// var db *sql.DB

// func CreateRoom(c *gin.Context) {
// 	var int_code int
// 	err := db.QueryRow("SELECT code FROM rooms WHERE is_available = true LIMIT 1").Scan(&int_code)

// 	// code := IntToBase36(int_code)
// 	if err == sql.ErrNoRows {
// 		log.Fatal("No available codes, please wait")
// 	} else if err != nil {
// 		fmt.Println("DatabaseError:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
// 		return
// 	} else {
// 		_, err = db.Exec("UPDATE rooms SET is_available = false WHERE code = $1", int_code)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room is_available", "desc": err})
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{"room_code": code})
// }

// func ReleaseRoom(c *gin.Context) {
// 	// err := services.ReleaseRoomService(c.Param("code"))
// 	if err == sql.ErrNoRows {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found to release"})
// 		return
// 	} else if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query room is_available"})
// 		return
// 	}

// 	// Update the room is_available to 'available'
// 	// err = services.UpdateStatusService()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to release room"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Room released successfully"})
// }
