// Example of map literals in golang

package main

import "fmt"

func main () {

	student := map[string] int {

		"age": 	 21, 
		"marks": 85,
	}
	fmt.Println(student)
}