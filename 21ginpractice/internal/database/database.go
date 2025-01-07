package internal

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("DB connected successfully!!")
	return db
}
