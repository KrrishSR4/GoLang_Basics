// Example of variadic functions in golang


package main

import "fmt"

func sum (numbers ...int) int {
	total := 0
	
	for _, number := range numbers {
		total += number
	}
	return total
}

func main () {

	result := sum (10, 20, 30, 40)

	fmt.Println(result)
}