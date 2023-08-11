//3. WAP in go language to accept user choice and print answer of using arithmetical operators.

package main

import "fmt"

func main() {
	var num1, num2, ch int

	fmt.Print("Enter 1st no. : ")
	fmt.Scan(&num1)

	fmt.Print("Enter 2nd no. : ")
	fmt.Scan(&num2)

	fmt.Print("\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n5. Remainder\nEnter Choice :")
	fmt.Scan(&ch)
	switch ch {
	case 1:
		fmt.Print("\n Addition :", num1+num2)
	case 2:
		fmt.Print("\n Subtraction :", num1-num2)
	case 3:
		fmt.Print("\n Multiplication : ", num1*num2)
	case 4:
		fmt.Print("\n Division : ", num1/num2)
	case 5:
		fmt.Print("\n Remainder : ", num1%num2)
	default:
		fmt.Print("\n Invalid Choice ...")
	}

	fmt.Println()
}