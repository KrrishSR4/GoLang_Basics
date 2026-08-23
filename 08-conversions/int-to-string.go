// Converting integer to string

package main

import (
	"fmt"
	"strconv"
)

func main () {
	var age int = 21

	var result = strconv.Itoa(age)

	fmt.Println(result)
	fmt.Printf("The type is: %T\n", result)

}