//3. Write a program in go language to copy all elements of one array into another using method.

package main

import "fmt"

func copy(x [100]int, n int) [100]int {
	var temp [100]int

	for i := 0; i < n; i++ {
		temp[i] = x[i]
	}

	return temp
}

func main() {
	var n int
	var arr [100]int
	var new [100]int

	fmt.Print("How many numbers you wants to enter : ")
	fmt.Scan(&n)

	fmt.Println("\nEnter Number : ")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	new = copy(arr, n)

	fmt.Println("\nCopied Array : ")
	for i := 0; i < n; i++ {
		fmt.Print(new[i], " ")
	}

	fmt.Println()
}
