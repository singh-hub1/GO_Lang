//1. Write a program in go language to create an interface and display it’s values with the help of type assertion.

package main

import "fmt"

func main() {
	var i interface{} = 5.2

	if v, result := i.(string)
	 result {
		fmt.Println("Value is : ", v, "\nIt is a String ")
	} else if v, result := i.(int)
	 result {
		fmt.Println("Value is : ", v, "\nIt is a Integer ")
	} else {
		v := i.(float64)
		fmt.Println("Value is : ", v, "\nIt is a Float")
	}

}

// Assertion is use for remove ambiguity
// It returns value and one boolean value, boolean value is use in if block. boolean value is true then execute if block.
// and also it help to reduce run time error (panic) when program fetch value from interface
