//1. WAP in go language to print addition of two number using function.

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var a, b int

	fmt.Print("Enter 1st no. : ")
	fmt.Scan(&a)

	fmt.Print("Enter 2nd no. : ")
	fmt.Scan(&b)

	fmt.Println("\nAddition of ", a, " and ", b, " is ", add(a, b))

}
