//3. WAP in go language to print Fibonacci series of n term

package main

import "fmt"

func main() {
	var a int = 0
	var b int = 1
	var c int = 0
	var n int

	fmt.Print("Enter a nth value : ")
	fmt.Scan(&n)

	fmt.Println()

	if n == 1 {
		fmt.Print(a)
	} else if n == 2 {
		fmt.Print(a, " ", b)
	} else {
		fmt.Print(a, " ")
		for i := 1; i < n; i++ {
			c = a + b
			fmt.Print(c, " ")
			a = b
			b = c
		}
	}

	fmt.Println()
}