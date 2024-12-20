package main

import "fmt"

func main() {
	var arr = [4]int{}
	fmt.Println(arr)
	fmt.Printf("type of Array: %T\n", arr) //type of int array

	var slice1 = []int{}
	fmt.Println("default slice values", slice1)
	fmt.Printf("type of Slice: %T\n", slice1) //type of int slice

	// slicing slices
	var slice2 = []int{1, 2, 3, 4, 5}
	slice2 = slice2[1:4]
	fmt.Println(slice2)

	// using make
	mslice := make([]int, 3)
	fmt.Println("slice using make:", mslice)

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	c[0] = "changed"
	fmt.Println("After modification:", c, s)

}
