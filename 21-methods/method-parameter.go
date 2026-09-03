// Example of parameters in methods in golang


package main

import "fmt"

type Student struct {
    Name string
}

func (s Student) greet(message string) {
    fmt.Println(s.Name, message)
}

func main() {
    student := Student{
        Name: "Rahul",
    }

    student.greet("Welcome!")
}