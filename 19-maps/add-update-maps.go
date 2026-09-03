// Example of adding and updating maps in golang


package main

import "fmt"

func main () {

	student := map[string] int {
		"age": 21,
	}

	student["marks"] = 85
	student["age"] = 22

	fmt.Println(student)
}