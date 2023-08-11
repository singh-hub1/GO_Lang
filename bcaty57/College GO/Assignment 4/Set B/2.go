//2. Write a program in go language to demonstrate working type switch in interface.

package main

import (
	"fmt"
	"strconv"
)

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("\n It is a String. ")
	case int:
		fmt.Println("\nIt is a Integer. ")
	case float64:
		fmt.Println("\nIt is a Float ")
	default:
		fmt.Println("Unknown Data Type ..")
	}
}

func main() {

	var ch interface{}
	var temp string

	fmt.Print("Enter anything : ")
	fmt.Scan(&temp)

	if v, err := strconv.Atoi(temp)
	      err == nil {
		ch = v
	} else if v, err := strconv.ParseFloat(temp,32)
	 err == nil {
		ch = v
	} else {
		ch = temp
	}

	findType(ch)
}
