package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	formattedTime := presentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println(formattedTime)

	createdDate := time.Date(2004, time.April, 24, 1, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
