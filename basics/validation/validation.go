package main

import (
	"fmt"
	"regexp"
)

func main() {
	email := "example@example.com"
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		fmt.Println("Email address is valid.")
	} else {
		fmt.Println("Email address is invalid.")
	}
}
