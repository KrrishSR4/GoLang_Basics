// Example of declaring a slice in golang


package main

import "fmt"

func main () {

	var numbers []int
	
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
