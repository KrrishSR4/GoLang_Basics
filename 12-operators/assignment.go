// Example of assignment operator in golang

package main

import "fmt"

func main () {

	a := 10

	a += 5
	fmt.Println(a)

	a -= 5
	fmt.Println(a)

	a *= 5
	fmt.Println(a)

	a /= 5
	fmt.Println(a)

	a %= 5
	fmt.Println(a)
}