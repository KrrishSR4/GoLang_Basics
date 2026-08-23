// Converting float64 to int

package main

import "fmt"

func main() {
	var price float64 = 99.99

	var result int = int(price)

	fmt.Println(result)
	fmt.Printf("The type is: %T\n", result)
}