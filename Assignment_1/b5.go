//Explain new function

package main

import "fmt"

func main() {
	a := new(int)
	//The new built-in function allocates memory. The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.
	Something(a)
	fmt.Println(*a)

}
func Something(a *int) {
	*a = 1
}
