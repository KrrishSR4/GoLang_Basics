// Example of an anonymous struct in golang


package main

import "fmt"

func main() {

    student := struct {

        Name string
        Age  int
    }{
		
        Name: "Rahul",
        Age:  21,
    }

    fmt.Println(student)
}