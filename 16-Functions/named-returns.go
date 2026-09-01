// Example of named returns in golang


package main

import "fmt"

func calculate (a int, b int) (sum int, product int) {

	sum = a + b
	product = a * b

	return
}

func main () {
	sum , product := calculate (10, 20)

	fmt.Println(sum)
	fmt.Println(product)
}