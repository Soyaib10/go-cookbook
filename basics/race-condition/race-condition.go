package main

import (
	"fmt"
	"sync"
)

func main() {
	var value int
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value++
		}()
	}

	wg.Wait()
	fmt.Println("Value:", value)
}
