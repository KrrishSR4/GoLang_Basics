// example of new() memory management in golang

package main 

import "fmt"

func main () {
	ptr := new (int)

	*ptr = 10

	fmt.Println(*ptr)
}