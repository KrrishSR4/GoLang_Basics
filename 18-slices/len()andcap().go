// Example of length and capacity of a slice in golang


package main

import "fmt"

func main () {

	numbers := [] int{10, 20, 30}

	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers, 40)
	
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}