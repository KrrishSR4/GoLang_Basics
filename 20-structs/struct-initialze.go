// Example of initializing a struct in golang


package main

import "fmt"

type student struct {

	Name	string
	Age		int
	Marks	float64
}

func main () {

	student := student {

		Name: "Krish",
		Age: 21,
		Marks: 85.5,
	}

	fmt.Println(student)
}

