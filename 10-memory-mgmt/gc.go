// example of garbage collection memory management in golang

package main 

import (
	"fmt"
	"runtime"
)

func main () {

	data := make ( [] int, 1000000)

	fmt.Println(len(data))

	data = nil

	runtime.GC()

	fmt.Println("Garbage collection completed")
}