package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex()
	channel()
}

func mutex() {
	var value int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			value++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Value:", value)
}

func channel() {
	const numGoroutines = 10

	dataCh := make(chan int, numGoroutines)
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dataCh <- 1
		}()
	}

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	total := 0
	for v := range dataCh {
		total += v
	}
	fmt.Println("Data value:", total)
}
