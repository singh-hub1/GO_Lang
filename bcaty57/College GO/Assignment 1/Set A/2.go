//2. WAP in go language to print whether number is even or odd.
package main

import "fmt"

func main() {
	var num int

	fmt.Printf("Enter a Nummer : ")
	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		fmt.Println("\nIt is a Even Number ...")
	} else {
		fmt.Println("\nIt is a Odd Number ....")
	}
}
