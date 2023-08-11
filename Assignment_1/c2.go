//WAP in go language to accept two strings and compare them.
package main

import (
	"fmt"
	"strings"
)

func main() {

	var s1 string
	var s2 string

	fmt.Println("Enter the first string :: ")
	fmt.Scanln(&s1)

	fmt.Println("Enter the second string :: ")
	fmt.Scanln(&s2)

	fmt.Println(strings.Compare(s1, s2))

}

/*
Returns 0 if the strings are equal (s1==s2)
Returns 1 if string 1 is greater than string 2 (s1 > s2)
Returns -1 is string 1 is lesser than string 2 (s1 < s2)
*/
