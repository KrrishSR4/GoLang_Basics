// Comma ok syntax in golang

package main 

import "fmt" 

func main () {
	ages := map[string]int{
		"Krish": 21,
	}

	age, ok := ages ["Krish"]

	fmt.Println(age)
	fmt.Println(ok)
}

