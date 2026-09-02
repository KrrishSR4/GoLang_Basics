// Example of creating a map with make() in golang


package main 

import "fmt"

func main () {

	student := make(map[string]int)

    student["age"] = 21
    student["marks"] = 85

    fmt.Println(student)
}