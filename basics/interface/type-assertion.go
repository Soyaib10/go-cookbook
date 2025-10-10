package main

import (
	"fmt"
)

// An interface and a type.
type Vehicle interface {
	Drive() string
}

type Car struct {
	Model string
}

func (c Car) Drive() string {
	return "All in!"
}

func checkVehicle(v Vehicle) {
	if car, ok := v.(Car); ok {
		fmt.Printf("Car model: %s\n", car.Model)
	} else {
		fmt.Println("Unknown type")
	}
}

func main() {
	var v Vehicle
	v = Car{Model: "Renault"}
	checkVehicle(v) // Output: Car model: Renault
}
