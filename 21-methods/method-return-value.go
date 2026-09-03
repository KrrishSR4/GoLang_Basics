// Example of methods in golang that returns a value


package main

import "fmt"

type Rectangle struct {
    Length float64
    Width  float64
}

func (r Rectangle) area() float64 {
    return r.Length * r.Width
}

func main() {
    rectangle := Rectangle{
        Length: 10,
        Width:  5,
    }

    result := rectangle.area()

    fmt.Println("Area:", result)
}