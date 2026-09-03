// Example of implicit interface in golang


package main

import "fmt"

type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof")
}

func main() {
    var s Speaker = Dog{}

    s.Speak()
}