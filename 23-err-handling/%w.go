package main

import (
	"errors"
	"fmt"
)

var ErrInvalidAge = errors.New("invalid age")

func validate(age int) error {
	if age < 18 {
		return fmt.Errorf("validation failed: %w", ErrInvalidAge)
	}

	return nil
}

func main() {
	err := validate(16)

	if err != nil {
		fmt.Println(err)
	}
}