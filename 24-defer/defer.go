// Example of defer keyword in golang


package main

import "fmt"

func main () {

	defer fmt.Println("This runs later")

	fmt.Println("Hello")
}