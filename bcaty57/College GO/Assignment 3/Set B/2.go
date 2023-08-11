//2. WAP in go language to sort array elements in ascending order.

package main

import "fmt"

func main() {
	var arr [10]int
	var n int

	fmt.Print("How many elemens you wants to enter : ")
	fmt.Scan(&n)

	fmt.Println("\nEnter Number :")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	fmt.Println("\nAfter ascending sort : ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// //This program is for slice using readymade library "sort"
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	// arr := []int{}

// 	var n int

// 	fmt.Print("How many numbers you wants to enter : ")
// 	fmt.Scan(&n)

// 	var arr = make([]int, n)

// 	fmt.Println("Enter Numbers : ")
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&arr[i])
// 	}

// 	sort.Ints(arr)

// 	fmt.Println("\nAfter ascending sort : ")
// 	for i := 0; i < n; i++ {
// 		fmt.Print(arr[i], " ")
// 	}
// 	fmt.Println()


// }
