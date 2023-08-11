//4. WAP in go language to illustrate pointer to pointer concept.

package main

import "fmt"

func main() {
	var a int = 10

	p := &a
	q := &p

	fmt.Println("Value of a is : ", **q)

	// first star for *p and second * for a
}
