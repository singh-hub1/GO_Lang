//WAP in go language to accept user choice and print answer of using
// arithmetical operators.

package main

import (
	"fmt"
)

func main() {
	var op string

	var number1, number2 float64

	fmt.Print("Please enter First number: ")
	fmt.Scanln(&number1)

	fmt.Print("Please enter Second number: ")
	fmt.Scanln(&number2)

	fmt.Print("Please enter Operator (+,-,/,%,*):")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Printf(" %0.2f %s %0.2f = %0.2f ", number1, op, number2, number1+number2)

	case "-":
		fmt.Printf(" %0.2f %s %0.2f = %0.2f ", number1, op, number2, number1-number2)

	case "*":
		fmt.Printf(" %0.2f %s %0.2f = %0.2f ", number1, op, number2, number1*number2)

	case "/":
		fmt.Printf(" %0.2f %s %0.2f = %0.2f ", number1, op, number2, number1/number2)

	default:
		fmt.Println("Invalid Operation")
	}

}
