package main

import (
	"fmt"
)

func main() {
	basic()
}

func basic() {
	ch := make(chan string)

	go func() {
		ch <- "Processing data from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}