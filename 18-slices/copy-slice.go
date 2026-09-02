// Example of copying a slice in golang


package main

import "fmt"

func main() {
    source := []int{10, 20, 30}
    destination := make([]int, 3)

    copy(destination, source)

    fmt.Println("Source:", source)
    fmt.Println("Destination:", destination)
}