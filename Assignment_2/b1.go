//1. WAP in go language to swap two numbers using call by reference concept.

package main

import "fmt"

func main() {

	var a int
	var b int

	fmt.Println("Enter First number:: ")
	fmt.Scanln(&a)

	fmt.Println("Enter Second number :: ")
	fmt.Scanln(&b)

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	swapReference(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)
}
func swapReference(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
