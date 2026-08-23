// Converting string to integer

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string = "100"

	result, err := strconv.Atoi(text)

	fmt.Println(result)
	fmt.Println(err)
}