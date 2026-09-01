// Example of an anonymous function in golang


package main 

import "fmt"

func main () {
	greet := func () {
		fmt.Println("Hello Krish")
	}

	greet ()
}