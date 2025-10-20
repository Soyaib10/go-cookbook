package main

import (
	"fmt"
	"time"
)

func worker(id int, work chan int, done chan bool) {
	for n := range work {
		fmt.Printf("Worker %d processing task %d\n", id, n)
		time.Sleep(time.Second)
		done <- true
	}
}

func main() {
	work := make(chan int, 5)
	done := make(chan bool, 5)

	for i := 1; i < 5; i++ {
		go worker(i, work, done)
	}

	for j := 0; j < 5; j++ {
		work <- j
	}
	close(work)

	for k := 0; k < 5; k++ {
		<-done
	}
	fmt.Println("All tasks processed")
}
