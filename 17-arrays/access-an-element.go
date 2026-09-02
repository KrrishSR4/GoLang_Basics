// Example of accessing an array's element in golang


package main

import "fmt"

func main () {

	numbers := [5] int{10, 20, 30, 40, 50}

	fmt.Println(numbers[4])
	fmt.Println(numbers[2])
	fmt.Println(numbers[0])
}