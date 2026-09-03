// Example of fmt.Errorf() dynamic error declaration in golang


package main

import (
	"fmt"
)

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is below the required age", age)
	}

	return nil
}

func main() {
	err := checkAge(16)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Valid age")
}