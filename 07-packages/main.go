// Using my own package

package main

import (
	"fmt"
	"07-packages/calculator"
)

func main() {
	result := calculator.Add(20, 30)

	fmt.Println("The sum is:", result)
}