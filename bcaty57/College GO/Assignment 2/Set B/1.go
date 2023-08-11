//1. WAP in go language to swap two numbers using call by reference concept.


package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Enter 1st no : ")
	fmt.Scan(&n1)

	fmt.Print("Enter 2nd no : ")
	fmt.Scan(&n2)

	swap(&n1, &n2)

	fmt.Println("\nAfter Swapping : ")
	fmt.Println("1st no. : ", n1)
	fmt.Println("2nd no. : ", n2)

}

func swap(x *int, y *int) {
	var temp int = *x
	*x = *y
	*y = temp
}