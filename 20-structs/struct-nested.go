// Example of a nested struct in golang


package main

import "fmt"

type Address struct {

    City  string
    State string
}

type Student struct {

    Name    string
    Age     int
    Address Address
}

func main() {

    student := Student{
		
        Name: "Rahul",
        Age:  21,
        Address: Address{
            City:  "Raipur",
            State: "Chhattisgarh",
        },
    }

    fmt.Println(student.Name)
    fmt.Println(student.Address.City)
    fmt.Println(student.Address.State)
}