// Example of defining an interface in golang


package main

import "fmt"

type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Dog says: Woof")
}

func main() {
    var speaker Speaker

    speaker = Dog{}

    speaker.Speak()
}