//WAP in go language to illustrate the concept of returning multiple values from
// a function

package main

import "fmt"

// myfunc return 3 values of int type
func myfunc(p, q int) (int, int, int) {
	return p - q, p * q, p + q
}

// Main Method
func main() {

	// The return values are assigned into
	// three different variables
	var (
		a int
		b int
	)

	fmt.Println("Enter first number ::")
	fmt.Scanln(&a)

	fmt.Println("Enter second number ::")
	fmt.Scanln(&b)

	var myVar1, myVar2, myVar3 = myfunc(a, b)
	fmt.Printf("\nResult of subtraction  is :: %d", myVar1)
	fmt.Printf("\nResult of multiplication :: %d", myVar2)
	fmt.Printf("\nResult of addition :: %d", myVar3)
}
