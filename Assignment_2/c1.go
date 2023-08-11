//WAP in go language to illustrate the concept of call by value.

package main

import "fmt"

func main() {

	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	swapValue(a, b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)
}
func swapValue(x, y int) {
	var temp int
	temp = x
	x = y
	y = temp
}
