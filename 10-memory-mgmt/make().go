// example of make memory management in golang

package main

import "fmt"

func main () {

	numbers := make ([] int, 5)

	numbers [0] = 10

	fmt.Println(numbers)
}