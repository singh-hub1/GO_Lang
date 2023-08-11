//WAP in go language to print recursive sum of digits of given number.
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter any number :: ")
	var a int
	fmt.Scanln(&a)
	res := sDigit(a)
	fmt.Println("Sum of digit of ", a, " are::  ", res)
}

func sDigit(a int) (int) {
	if a == 0 {
		return 0
	}
	return(a%10 + sDigit(a/10))
}
