//4. WAP in go language to check whether accepted number is single digit or not.

package main

import "fmt"

func main() {
	var a int

	fmt.Print("Enter a number :  ")
	fmt.Scan(&a)

	if a/10 == 0 {
		fmt.Println("\nsingle digit ")
	} else {
		fmt.Println("\nNOT single digit ")
	}
}
