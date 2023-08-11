//WAP in go language to demonstrate use of names returns variables

package main

import "fmt"

func main() {

	var a int
	var b int

	fmt.Println("Enter first number :: ")
	fmt.Scanln(&a)

	fmt.Println("Enter Second number :: ")
	fmt.Scanln(&b)

	m, d := calculator(a, b)
	fmt.Println("105 x 7 = ", m)
	fmt.Println("105 / 7 = ", d)
}

// function having named arguments
//return types is same

func calculator(a, b int) (mul int, div int) {
	mul = a * b
	div = a / b

	return
}
