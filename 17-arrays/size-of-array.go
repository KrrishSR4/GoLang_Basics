// Example of determining size of an array in golang


package main

import "fmt"

func main () {

	numbers := [...] int{10, 20, 30, 40, 50}

	fmt.Println(numbers)
	fmt.Println(len(numbers))
}