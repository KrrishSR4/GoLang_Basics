// exmple -2
package main

import "fmt"

var language = "Go"

func main() {
	age := 21

	if age >= 18 {
		message := "Adult"
		fmt.Println(message)
	}

	fmt.Println(language)

}

// prints "Adult" and "Go"
