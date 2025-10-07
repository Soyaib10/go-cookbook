package main

import (
	"fmt"
	"slices"
)

func main() {
	// Create a slice with length 3 and capacity 10.
	slice := make([]int, 3, 10)
	fmt.Println(slice) // Output: [0 0 0].

	slicing()
	appending()
	sliceUtility()
}

func slicing() {
	arr := [5]int{1, 2, 3, 4, 5}
	// Create a slice from the array.
	slice := arr[1:4]
	fmt.Println(slice) // Output: [2 3 4].
}

func appending() {
	// Initial slice.
	slice := []int{1, 2, 3}
	// Append one element.
	slice = append(slice, 4)
	// Append multiple elements.
	slice = append(slice, 5, 6, 7)
	fmt.Println(slice) // Output: [1 2 3 4 5 6 7].
}

func sliceUtility() {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// Find elements
	fmt.Printf("Contains 5: %t\n", slices.Contains(numbers, 5))
	fmt.Printf("Index of 4: %d\n", slices.Index(numbers, 4))

	// Sorting and binary search
	sortedNumbers := slices.Clone(numbers)
	slices.Sort(sortedNumbers)
	fmt.Printf("Sorted: %v\n", sortedNumbers)
	binarySearchRes, _ := slices.BinarySearch(sortedNumbers, 5)
	fmt.Printf("Binary search for 5: %d\n", binarySearchRes)

	// Insert and delete
	inserted := slices.Insert(numbers, 2, 99, 100)
	fmt.Printf("After insert: %v\n", inserted)

	deleted := slices.Delete(numbers, 1, 3)
	fmt.Printf("After delete: %v\n", deleted)

	// Comparison and equality
	other := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("Equal to other: %t\n", slices.Equal(numbers, other))

	// Reverse
	reversed := slices.Clone(numbers)
	slices.Reverse(reversed)
	fmt.Printf("Reversed: %v\n", reversed)
}
