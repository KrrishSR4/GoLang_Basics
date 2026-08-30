// example of stack memory management in golang

package main 

import "fmt"

func add (a int, b int) int {

	x := a
	y := b

	return x + y
}

func main () {
	result := add (10, 20)

	fmt.Println(result)
}