// Example of function that accepts an interface in golang


package main

import "fmt"

type Speaker interface {
    Speak()
}

type Dog struct{}

type Cat struct{}

func (d Dog) Speak() {
    fmt.Println("Dog: Woof")
}

func (c Cat) Speak() {
    fmt.Println("Cat: Meow")
}

func makeSpeak(s Speaker) {
    s.Speak()
}

func main() {
    dog := Dog{}
    cat := Cat{}

    makeSpeak(dog)
    makeSpeak(cat)
}