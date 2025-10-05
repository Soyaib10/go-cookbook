package main

import "fmt"

func main() {
	// Create a slice with length 3 and capacity 10.
	slice := make([]int, 3, 10)
	fmt.Println(slice) // Output: [0 0 0].

	slicing()
}

func slicing() {	
	arr := [5]int{1, 2, 3, 4, 5}
	// Create a slice from the array.
	slice := arr[1:4]
	fmt.Println(slice) // Output: [2 3 4].
}
