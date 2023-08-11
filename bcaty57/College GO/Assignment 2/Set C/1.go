//1. WAP in go language to illustrate the concept of call by value.

package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Enter 1st no : ")
	fmt.Scan(&n1)

	fmt.Print("Enter 2nd no : ")
	fmt.Scan(&n2)

	// Call by Value
	swap(n1, n2)

	fmt.Println("\nAfter swapping (Call by value method) ")
	fmt.Println("1st no : ",n1)
	fmt.Println("2nd no : ",n2)
}

func swap(n1, n2 int) {
	temp := n1
	n1 = n2
	n2 = temp
}