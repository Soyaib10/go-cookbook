package main

import (
	"fmt"
)

func main() {
	var x int = 1
	var p *int = &x // Pointer to x.

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p)
	fmt.Println("Value via pointer p:", *p)

	// Modifying the value of x using a pointer.
	*p = 2
	fmt.Println("New value of x:", x)
}
