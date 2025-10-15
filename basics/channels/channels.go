package main

import (
	"fmt"
	"time"
)

func main() {
	basic()
	buffer()
	closingChannel()
	selectStatement()
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

func selectStatement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Data from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Data from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
