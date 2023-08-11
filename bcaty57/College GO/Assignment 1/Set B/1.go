//1. WAP in go to print table of given number.

package main

import "fmt"

func main() {
	var num int

	fmt.Printf("Enter a Number : ")
	fmt.Scanf("%d", &num)

	fmt.Println()
	for i := 1; i <= 10; i++ {
		fmt.Println(i, " X ", num, " = ", i*num)

	}

}