package main

import "fmt"

type student struct {
	name     string
	age      int
	course   string
	studying string
}

//copy of actual struct instance is passes
func (s student) printName() {
	fmt.Println("Name of the student is:", s.name)
}

//struct instance pointer is passed, i.e, it sets the name of actual struct instance that called set method
func (s *student) setName(newName string) {
	s.name = newName
}

func main() {
	student := student{
		name:     "John",
		age:      20,
		course:   "Computer Science",
		studying: "final year",
	}
	student.printName()

	//trying to change name by not using pointers
	student.setName("Peter")
	student.printName()
}
