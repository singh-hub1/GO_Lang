//2.WAP in go language to accept two strings and compare them.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Print("Enter a 1st string : ")
	fmt.Scan(&str1)

	fmt.Print("Enter a 2nd string : ")
	fmt.Scan(&str2)

	result := strings.Compare(str1, str2)

	fmt.Println()
	if -1 == result {
		fmt.Println("2nd String is Greater...")
	} else if 1 == result {
		fmt.Println("1nd String is Greater ")
	} else {
		fmt.Println("Both are equal")
	}
}