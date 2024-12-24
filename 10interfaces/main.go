package main

import "fmt"

// an interface type is a set of method signatures
type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func (c circle) area() float64 {
	return (3.14 * c.radius * c.radius)
}

func (c circle) perimeter() float64 {
	return (2 * 3.14 * c.radius)
}

func (r rect) area() float64 {
	return (r.width * r.height)
}

func (r rect) perimeter() float64 {
	return (2 * (r.width + r.height))
}

func calc(g geometry) {
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perimeter())
}

func main() {
	c := circle{radius: 5.0}
	r := rect{width: 4.0, height: 6.0}

	fmt.Println("Area & perimeter of circle: ")
	calc(c)
	fmt.Println("\n\nArea & perimeter if rectangle: ")
	calc(r)

	// Type switches
	do(21)
	do("hello")
	do(true)

}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
