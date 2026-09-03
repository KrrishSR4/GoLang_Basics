// Example of determining length of a map in golang


package main

import "fmt"

func main() { 

    student := map[string]int{
		
        "age":   21,
        "marks": 85,
        "year":  2026,
    }

    fmt.Println(len(student))
}