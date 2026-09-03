// Example of creating an error in golang


package main 

import (
	"fmt"
	"errors"
)

func main () {

	err := errors.New("Something went wrong")

	fmt.Println(err)
}