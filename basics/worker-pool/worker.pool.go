package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing task %d\n", id, j)
		// Simulate work by multiplying the number by 2.
		results <- j * 2
		fmt.Printf("Worker %d completed task %d\n", id, j)
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers.
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to the workers.
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to finish.
	wg.Wait()
	close(results)

	// Collect results.
	for result := range results {
		fmt.Println("Result:", result)
	}
}
