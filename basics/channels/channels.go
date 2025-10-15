package main

import (
	"fmt"
)

func main() {
	basic()
	buffer()
	closingChannel()
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


func closingChannel() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
