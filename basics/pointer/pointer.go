package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	var x int = 1
	var p *int = &x // Pointer to x.

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p)
	fmt.Println("Value via pointer p:", *p)

	// Modifying the value of x using a pointer.
	*p = 2
	fmt.Println("New value of x:", x)

	pointerstruct()
}

func pointerstruct() {
	p := Person{"Alice", 20}
	ptr := &p

	fmt.Println("Name:", ptr.name)
	fmt.Println("Age:", ptr.age)

	// Updating struct fields via pointer.
	ptr.age = 25
	fmt.Println("Updated Age:", ptr.age)
}