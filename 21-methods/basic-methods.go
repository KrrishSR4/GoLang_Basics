// Example of basic methods in golang


package main

import "fmt"

type Student struct {

    Name string
    Age  int
}

func (s Student) display() {

    fmt.Println("Name:", s.Name)
    fmt.Println("Age:", s.Age)
}

func main() {

    student := Student{
        Name: "Rahul",
        Age:  21,
    }

    student.display()
}