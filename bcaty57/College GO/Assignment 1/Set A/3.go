//3. WAP in go language to swap the number without temporary variable.

package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Enter a 1st no. : ")
	fmt.Scanf("%d", &a)

	fmt.Print("Enter a 2nd no. : ")
	fmt.Scanf("%d", &b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	/*
	a = a+b
	b = a-b
	a = a-b
	*/

	/*
	a = a*b
	b = a/b
	a = a/b
	*/



	fmt.Println("\n1st no. ", a, "\n2nd no. ", b)
}