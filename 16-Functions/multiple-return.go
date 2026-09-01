// Example of Multiple return values in golang

package main

import "fmt"

func calculate (a int, b int) (int, int) {
	return a + b, a * b
}

func main () {
	sum, product := calculate(10, 20)

	fmt.Println(sum)
	fmt.Println(product)
}