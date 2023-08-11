//. WAP in go language to print addition of two number using function.
package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter fisrt number ::  ")
	var a, b int
	fmt.Scanln(&a)

	fmt.Println("Enter second number ::  ")
	fmt.Scanln(&b)

	res := sum(a, b)
	fmt.Println("Additon of two numbers :: ", res)

}

func sum(x, y int) int {
	return (x + y)
}
