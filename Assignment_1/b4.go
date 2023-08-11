//poiter to pointer concept
package main

import "fmt"

func main() {
	var a int=5
	var x *int
	var y **int

	x = &a
	y = &x

	fmt.Println("The value of x is %d::", *x)
	fmt.Println("The value of y is %d::", **y)
}
