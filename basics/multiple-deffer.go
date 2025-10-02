package main

import "fmt"

func MultipleDefer() {
	defer fmt.Println("Bob")
	defer fmt.Println("Alex")

	fmt.Println("Welcome")
}
