// converting integer to float

package main
 
import "fmt"

func main () {
	var num int = 20

	var result float64 = float64(num)

	fmt.Println(result)
	fmt.Printf("The type is: %T\n", result )
}