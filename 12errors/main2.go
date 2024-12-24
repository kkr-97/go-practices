package main

import (
	"fmt"
	"math"
)

// custom error
type ErrSqrt struct {
	num int
}

func (e ErrSqrt) Error() string {
	return fmt.Sprintf("square root cannot be found for %v, since it is negative", int(e.num))
}

func Sqrt2(num int) (float64, error) {
	if num < 0 {
		return 0, ErrSqrt{num}
	}
	return float64(math.Sqrt(float64(num))), nil
}

func CustomErrorMain() {
	num := -2
	//create a custom error
	result, err := Sqrt2(num)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Custom error example: ", result)
	}

}
