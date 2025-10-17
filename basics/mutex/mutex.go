package main

import (
	"fmt"
	"sync"
)

func main() {
	mapWithMutex()
	RWMutex()
}

func mapWithMutex() {
	var mu sync.Mutex
	m := make(map[string]int)

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[fmt.Sprintf("task%d", i)] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	mu.Lock()
	fmt.Println("Data contents:", m)
	mu.Unlock()
}


func RWMutex() {
	var rwMu sync.RWMutex
	m := make(map[string]int)

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rwMu.Lock()
			m[fmt.Sprintf("task%d", i)] = i
			rwMu.Unlock()
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		rwMu.RLock()
		fmt.Println("Reading data contents:", m)
		rwMu.RUnlock()
	}()

	wg.Wait()
}

