package main

import (
	"fmt"
)

type Number interface {
	int | int64 | float32 | float64 | int8
}

func add[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(2, 3))
}
