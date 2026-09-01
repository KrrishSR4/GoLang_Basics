// Example of infinite loop in golang

package main 

import "fmt"

func main () {

	for {
		fmt.Println("Running")

		break
	}
}