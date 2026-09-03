// Example of loop through map in golang


package main

import "fmt"

func main() {

    student := map[string]int{

        "age":   21,
        "marks": 85,
        "year":  2026,
    }

    for key, value := range student {
		
        fmt.Println(key, value)
    }
}