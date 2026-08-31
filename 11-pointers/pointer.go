// Example of using pointer in golang


package main 

import "fmt"

func main () {

	age := 21
	ptr := &age

	fmt.Println(age)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}