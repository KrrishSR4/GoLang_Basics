// Example of high - order functions in golang


package main

import "fmt"

func calculate(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	result := calculate(10, 20, add)

	fmt.Println(result)
}