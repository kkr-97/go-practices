package main

import (
	"errors"
	"fmt"
	"math"
)

// using errors package
var ErrNegSqrt = errors.New("number cannot be negative")

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrNegSqrt
	}
	return float64(math.Sqrt(num)), nil
}

func main() {
	//using errors pkg
	result, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Square root of 16 is:", result)
	}

	//using custom error
	CustomErrorMain()
}
