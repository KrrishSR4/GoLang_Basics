// Example of append() in slices


package main

import "fmt"

func main () {

	numbers := [] int{10, 20, 30}

	numbers = append(numbers, 40)

	fmt.Println(numbers)
}

// append() is used to add an element in the slice