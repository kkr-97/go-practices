package internal

type Note struct {
	Id     int    `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
