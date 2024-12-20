package main

import "fmt"

func main() {
	fmt.Println("Arrays: ")

	var numsArr [3]int // default value is 0
	numsArr[0] = 10
	numsArr[1] = 20
	fmt.Println("Given Array: ", numsArr)
	fmt.Println("lenght:", len(numsArr))

	var fltArr = [3]float64{1.1, 2.2}
	fmt.Println("Given Array: ", fltArr)

	fmt.Printf("type of 0 %T", 0.0)
}
