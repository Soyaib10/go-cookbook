package main

import (
	"flag"
	"fmt"
	"log"	
)

func main() {
	task := flag.String("task", "", "Task description.")	
	priority := flag.Int("priority", 1, "Priority (1-5).")
	flag.Parse()

	if *task == "" {
		fmt.Println("Error: Task required.")
		log.Fatal("task parameter is required")
	}

	fmt.Printf("Added task: %s (Priority: %d)\n", *task, *priority)
}
