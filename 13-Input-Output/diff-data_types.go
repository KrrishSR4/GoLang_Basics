// Example of taking different data types as input in golang


package main 

import "fmt"

func main () {

	var latency int
	var slo float64
	var uptime bool

	fmt.Print("What's the latency: ")
	fmt.Scan(&latency)

	fmt.Print("What's the slo: ")
	fmt.Scan(&slo)

	fmt.Print("Is website down: ")
	fmt.Scan(&uptime)

	fmt.Println("Latency: ", latency)
	fmt.Println("Slo: ", slo)
	fmt.Println("Uptime: ", uptime)

	
}