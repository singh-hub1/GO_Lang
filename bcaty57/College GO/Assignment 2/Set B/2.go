//2. WAP in go language to demonstrate use of names returns variables.

package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Enter 1st no : ")
	fmt.Scan(&n1)

	fmt.Print("Enter 2nd no :")
	fmt.Scan(&n2)

	add, sub, mul, div := arith(n1, n2)

	fmt.Println("\nAddition : ", add)
	fmt.Println("Subtraction : ", sub)
	fmt.Println("Multiplication : ", mul)
	fmt.Println("Division : ", div)
}

func arith(n1, n2 int) (add, sub, mul, div int) {
	add = n1 + n2
	sub = n1 - n2
	mul = n1 * n2
	div = n1 / n2

	return
}