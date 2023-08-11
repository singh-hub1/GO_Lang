//1. WAP in go language to find the largest and smallest number in an array.

package main

import "fmt"

func main() {
	var n int
	var arr [100]int

	fmt.Print("How Many elements you wants to enter : ")
	fmt.Scan(&n)

	fmt.Println("\nEnter numbers : ")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	min := arr[0]
	for i := 1; i < n; i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	fmt.Println("\nSmallest number in given number : ", min)

}
