// panic is for situations where the program cannot safely continue.


package main

import "fmt"

func main() {
	fmt.Println("Starting")

	panic("something went seriously wrong")
}