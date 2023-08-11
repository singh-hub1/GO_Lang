//3. WAP in go language to illustrate the concept of returning multiple values from a function

package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Enter 1st no : ")
	fmt.Scan(&n1)

	fmt.Print("Enter 2nd no : ")
	fmt.Scan(&n2)

	add, sub, mul, div := arith(n1, n2)

	fmt.Println("\nAddition : ", add)
	fmt.Println("Subtraction : ", sub)
	fmt.Println("Multiplication : ", mul)
	fmt.Println("Division : ", div)

}

func arith(n1, n2 int) (int, int, int, int) {
	add := n1 + n2
	sub := n1 - n2
	mul := n1 * n2
	div := n1 / n2

	return add, sub, mul, div

}