//WAP in go language to swap the number without temporary variable.
package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Println("Enter the first number :: ")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number :: ")
	fmt.Scanln(&b)

	fmt.Println("Before swapping the number a=", a, " and b=", b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("After swapping the number a=", a, " and b=", b)

}
