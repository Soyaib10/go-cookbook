package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `validate:"required"`
	Age   int    `validate:"gte=0,lte=130"`
	Email string `validate:"required,email"`
}

func main() {
	v := validator.New()

	user := User{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}

	err := v.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Fatal(err)
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Validation Error:", err)
			log.Fatal(err)
		}
	} else {
		fmt.Println("User input is valid!")
	}
}
