package main

import (
	"fmt"
)

func main() {
	// Create a map with string keys and int values.
	ages := map[string]int{
		"Alice": 15,
		"Bob":   20,
	}

	// Add a key-value element.
	ages["Chris"] = 25

	// Retrieve and print a value.
	if age, exists := ages["Alice"]; exists {
		fmt.Printf("Alice is %d years old.\n", age)
	} else {
		fmt.Println("Age for Alice not found.")
	}
	
	// Remove a key-value element.
	delete(ages, "Bob")

	// Print the map.
	fmt.Println("Updated ages:", ages)
	Key()
	nestedmaps()
}

func Key() {
	// Initialize a map.
	cities := map[string]string{
		"NY": "New York",
		"LA": "Los Angeles",
	}

	// Check if a key exists.
	if city, found := cities["NY"]; found {
		fmt.Printf("The abbreviation &#39;NY&#39; stands for %s.\n", city)
	} else {
		fmt.Println("Key not found.")
	}

	// Iterate over a map.
	for abbrev, city := range cities {
		fmt.Printf("%s -> %s\n", abbrev, city)
	}
}

func nestedmaps() {
	departments := make(map[string]map[string]int)

	// Add nested maps.
	departments["HR"] = map[string]int{
		"Alice": 35,
		"Bob":   40,
	}

	// Retrieve a nested map.
	if hrDept, exists := departments["HR"]; exists {
		for name, age := range hrDept {
			fmt.Printf("%s in HR is %d years old.\n", name, age)
		}
	}
}
