//WAP in go language to print whether number is even or odd
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the any number :: ")
	var a int
	fmt.Scanln(&a)

	for i := 0; i <= a; i++ {
		if i%2 == 0 {
			fmt.Println("Number::", i, " is even.")
		} else {
			fmt.Println("Number::", i, "is odd. ")
		}
	}

}
