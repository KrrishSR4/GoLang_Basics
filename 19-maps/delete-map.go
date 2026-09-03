// Example of deleting elements from a map in golang


package main 

import "fmt"

func main () {

	student := map[string] int {

		"age": 21,
		"marks": 85,
	}

	delete(student, "age")
	fmt.Println(student)
}