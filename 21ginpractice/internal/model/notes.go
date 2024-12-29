package internal

type Note struct {
	Id     int    `gorm:"primaryKey"`
	title  string `json:"title"`
	author string `json:"title"`
}
