package main

import (
	"errors"
	"fmt"
	"math"
)

// using errors package
var ErrNegSqrt = errors.New("Number Cannot Be Negative")

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrNegSqrt
	}
	return float64(math.Sqrt(num)), nil
}

func main() {
	result, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Square root of 16 is:", result)
	}
}
