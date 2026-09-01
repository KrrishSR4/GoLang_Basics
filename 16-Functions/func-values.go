// Example of function values in golang


package main

import "fmt"

func add (a int, b int) int {
	return a + b
}

func main () {

	operation := add

	result := operation(500, 100)

	fmt.Println(result)
}