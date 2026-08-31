// Example of if - else - if in golang

package main

import "fmt"

func main () {

	marks := 9

	if marks >= 60 {
		fmt.Println("A")
	} else if marks >= 45 {
		fmt.Println("Just pass")
	} else {
		fmt.Println("Please work hard you're fail")
	}
}