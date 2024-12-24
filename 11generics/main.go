package main

import "fmt"

func FindIndex[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Test the function with a slice of integers
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("index of 9: %d", FindIndex(intSlice, 9))

	// Test the function with a slice of strings
	stringSlice := []string{"apple", "banana", "cherry", "date", "elder"}
	fmt.Printf("index of 'grape': %d", FindIndex(stringSlice, "date"))

}
