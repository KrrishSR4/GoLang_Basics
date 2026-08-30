// example of escape analysis memory management in golang

package main

import "fmt"

func getPointer () *int {

	x:= 10
	return &x
}

func main () {

	ptr := getPointer ()
	fmt.Println(*ptr)
}