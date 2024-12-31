package main

import "fmt"

func main() {
	var ptr1 *int
	var ptr2 *int
	fmt.Println(ptr1)

	num := 24
	ptr2 = &num
	fmt.Println("reference", ptr2)
	fmt.Println("value", *ptr2)

	*ptr2++
	fmt.Println("value after increment", *ptr2)
}
