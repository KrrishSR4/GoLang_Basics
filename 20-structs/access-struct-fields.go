// Example of accessing struct fields using . operator


package main

import "fmt"

type Student struct {

    Name  string
    Age   int
    Marks float64
}

func main() {

    student := Student{
		
        Name:  "Rahul",
        Age:   21,
        Marks: 85.5,
    }

    fmt.Println(student.Name)
    fmt.Println(student.Age)
    fmt.Println(student.Marks)
}