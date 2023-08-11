//WAP in go language using function to check whether accepts number is
// palindrome or not.
package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter any number :: ")
	fmt.Scanln(&a)

	var result int
	result = pallindrome(a)

	if a == result {
		fmt.Println("The number ", a, " is Pallindrome.")
	} else {
		fmt.Println("The number ", a, " is not  Pallindrome.")
	}

}
func pallindrome(a int) (int) {
	var r, s int

	for ; a > 0; a = a / 10 {
		r = a % 10
		s = s*10 + r
	}

	return s
}
