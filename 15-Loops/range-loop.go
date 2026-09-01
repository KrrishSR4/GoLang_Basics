// Example of range loop in golang

package main

import "fmt"

func main () {

	numbers := [] int {10, 20, 30}

	for index, value := range numbers {
		fmt.Println(index, value)
	}
}