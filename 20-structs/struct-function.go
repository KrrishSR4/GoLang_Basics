// Example of a function passed over a struct in golang


package main

import "fmt"

type Student struct {

    Name string
    Age  int
}

func printStudent(student Student) {

    fmt.Println(student.Name)
    fmt.Println(student.Age)
}

func main() {

    student := Student{

        Name: "Rahul",
        Age:  21,
    }

    printStudent(student)
}