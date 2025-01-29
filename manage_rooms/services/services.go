package services

// import (
// 	"database/sql"
// 	"fmt"
// 	"math"
// 	"strings"
// )

// type ManageRoomsService struct {
// 	db *sql.DB
// }

// func (s *ManageRoomsService) InitServices(dbInstance *sql.DB) {
// 	s.db = dbInstance
// }

// const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// func Base36ToInt(code string) int {
// 	code = strings.ToUpper(code) // Ensure it's uppercase
// 	value := 0
// 	for i, char := range code {
// 		pos := strings.IndexRune(charset, char) // Find position in charset
// 		if pos == -1 {
// 			panic(fmt.Sprintf("Invalid character: %c", char))
// 		}
// 		exp := 3 - i
// 		value += pos * int(math.Pow(36, float64(exp)))
// 	}
// 	return value
// }

// // IntToBase36 converts an integer to a Base-36 string
// func IntToBase36(num int) string {
// 	if num < 0 {
// 		panic("Negative numbers are not allowed")
// 	}
// 	result := ""
// 	for num > 0 {
// 		ind := num % 36
// 		result = string(charset[ind]) + result
// 		num /= 36
// 	}
// 	for len(result) < 4 {
// 		result = "0" + result
// 	}
// 	return result
// }

// func (s *ManageRoomsService) CreateRoomService() (string, error) {
// 	var int_code int
// 	err := s.db.QueryRow("SELECT code FROM rooms WHERE is_available = true LIMIT 1").Scan(&int_code)

// 	code := IntToBase36(int_code)
// 	if err == sql.ErrNoRows {
// 		return "", err
// 	} else if err != nil {
// 		fmt.Println("DatabaseError:", err)
// 		return "", err
// 	} else {
// 		_, err = s.db.Exec("UPDATE rooms SET is_available = false WHERE code = $1", int_code)
// 		if err != nil {
// 			return "", err
// 		}
// 	}
// 	return code, nil

// }

// func (s *ManageRoomsService) ReleaseRoomService(codeS string) error {
// 	code := Base36ToInt(codeS)

// 	// Check if the room exists and is not already released
// 	var is_available bool
// 	err := s.db.QueryRow("SELECT is_available FROM rooms WHERE code = $1", code).Scan(&is_available)
// 	if err == sql.ErrNoRows {
// 		return err
// 	} else if err != nil {
// 		return err
// 	}

// 	if is_available {
// 		return nil
// 	}

// 	// Update the room is_available to 'available'
// 	_, err = s.db.Exec("UPDATE rooms SET is_available = true WHERE code = $1", code)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
