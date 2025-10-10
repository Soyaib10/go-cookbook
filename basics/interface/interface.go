package main

import (
	"fmt"
)

// Define an interface.
type Greeter interface {
	Greet() string
}

// Define a type that satisfies our interface.
type Human struct {
	Name string
}

func (h Human) Greet() string {
	return "Hello!"
}

// Another type that satisfies the interface.
type Cow struct {
	Name string
}

func (c Cow) Greet() string {
	return "Moo?"
}

func main() {
	var g Greeter
	human := Human{Name: "Alex"}
	cow := Cow{Name: "Milka"}

	g = human
	fmt.Println(g.Greet()) // Output: Hello!

	g = cow
	fmt.Println(g.Greet()) // Output: Moo!
}
