package main

import "fmt"

func process() {
	defer cleanup()

	fmt.Println("Processing...")
	// Simulate some processing.
}

func cleanup() {
	fmt.Println("Cleaning up resources.")
}

