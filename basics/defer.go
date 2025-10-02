package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString("File content")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File written successfully")

	MultipleDefer()
	process()
}
