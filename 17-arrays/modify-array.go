// Example of modifying an array in golang


package main

import "fmt"

func main () {

	numbers := [5] int {10, 20, 30, 40, 50}

	numbers[1] = 100

	fmt.Println(numbers)
}