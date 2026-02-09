package main

import "fmt"

type Pair[T any, U any] struct {
	First  T
	Second U
}

func Length[T any](s []T) int {
	return len(s)
}

func main() {
	basic()
	structParam()
}

func basic() {
	ints := []int{1, 2, 3, 4}
	strs := []string{"a", "b", "c"}

	fmt.Println("Length of ints:", Length(ints))
	fmt.Println("Length of strs:", Length(strs))
}

func structParam() {
	intStr := Pair[int, string]{1, "one"}
	fmt.Printf("Pair: %v\n", intStr)

	strStr := Pair[string, string]{"hello", "world"}
	fmt.Printf("Pair: %v\n", strStr)
}
