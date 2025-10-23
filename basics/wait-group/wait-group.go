package main

import (
	"fmt"
	"sync"
)

const numWorkers = 5

func worker(id int, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()

	// Simulate a task.
	if id%2 == 0 {
		errChan <- fmt.Errorf("task %d encountered an error", id)
		return
	}

	fmt.Printf("Task %d completed successfully\n", id)
}

func main() {
	var wg sync.WaitGroup
	errChan := make(chan error, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, errChan)
	}


	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("All tasks completed")
}
