//1. WAP in go language to concatenate two strings using pointers.

package main

import "fmt"


func main() {
	var str1, str2 string

	fmt.Print("Enter a 1st String : ")
	fmt.Scan("%s", &str1)

	fmt.Print("Enter a 2nd String : ")
	fmt.Scan("%s", &str2)

	ptrStr1 := &str1
	ptrStr2 := &str2

	final := *ptrStr1 + *ptrStr2

	fmt.Println("\nAfter concatenation : ", final)
}
