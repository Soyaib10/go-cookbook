package main

import (
	"fmt"
)

func main() {
	basic()
	buffer()
}

func basic() {
	ch := make(chan string)

	go func() {
		ch <- "Processing data from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}

func buffer() {
	ch := make(chan int, 2)

	ch <- 42
	ch <- 73

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}