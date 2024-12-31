package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	formattedTime := presentTime.Format("01") //month
	fmt.Println("Month: ", formattedTime)
	formattedTime = presentTime.Format("02") //date
	fmt.Println("Date: ", formattedTime)
	formattedTime = presentTime.Format("2006") //year
	fmt.Println("Year: ", formattedTime)

	createdDate := time.Date(2004, time.April, 24, 1, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
