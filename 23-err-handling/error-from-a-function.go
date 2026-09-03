// Example of returning an error from a function 


package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) (string, error) {
	if age < 18 {
		return "", errors.New("age must be 18 or above")
	}

	return "Access granted", nil
}

func main() {
	result, err := checkAge(10)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}