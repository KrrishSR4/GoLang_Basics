// Example of slice from an array in golang


package main

import "fmt"

func main () {

	numbers := [5] int{10, 20, 30, 40, 50}

	slice := numbers[1:4]

	fmt.Println(slice)
}