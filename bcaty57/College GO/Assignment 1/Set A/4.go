//4. WAP in go Language to print address of a variable.

package main

import "fmt"

func main() {
	var a int

	fmt.Printf("Address is %p\n", &a)
}
