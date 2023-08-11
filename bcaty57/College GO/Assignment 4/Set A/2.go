//2. Write a program in go language to print multiplication of two numbers using method.

package main

import "fmt"

func main() {
	var n1 int
	var n2 int

	fmt.Print("Enter 1st number : ")
	fmt.Scan(&n1)
	fmt.Print("Enter 2nd number : ")
	fmt.Scan(&n2)

	fmt.Println("\nMultiplication of ", n1, " and ", n2, " is ", multiplication(n1, n2))
}

func multiplication(x, y int) int {
	return x * y
}
