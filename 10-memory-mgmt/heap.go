// example of heap memory managemant in golang

package main 

import "fmt"

func createUser () *int {

	age := 21
	return &age
}

func main () {
	age := createUser ()
	fmt.Println(*age)
}