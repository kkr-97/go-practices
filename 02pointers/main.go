package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println(ptr)

	num := 24
	var ptr = &num
	fmt.Println("reference", ptr)
	fmt.Println("value", *ptr)

	*ptr++
	fmt.Println("value after increment", *ptr)
}
