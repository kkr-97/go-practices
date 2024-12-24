package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrNegSqrt = errors.New("Number cannot be negative")

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrNegSqrt
	}
	return float64(math.Sqrt(num)), nil
}

func main() {
	// Test case 1: Positive number
	result, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Square root of 16 is:", result)
	}
}
