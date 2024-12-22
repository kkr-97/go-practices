package main

import "fmt"

type student struct {
	name     string
	age      int
	course   string
	studying string
}

func main() {
	student := student{
		name:     "John",
		age:      20,
		course:   "Computer Science",
		studying: "final year",
	}
	fmt.Println("student details:", student)
	fmt.Printf("detailed view: %+v", student)
}
