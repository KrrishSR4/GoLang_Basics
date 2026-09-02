// Example of make()  with slices in golang


package main

import "fmt"

func main () {

	numbers := make([]int, 3)

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	
	fmt.Println(numbers)
}