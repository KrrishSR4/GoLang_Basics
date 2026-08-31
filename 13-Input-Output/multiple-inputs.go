// Example of Multiple inputs in golang

package main 

import "fmt"

func main () {

	var name string
	var age int
	
	fmt.Print("Enter name and age: ")
	fmt.Scan(&name, &age)

	fmt.Println(name, age)
}