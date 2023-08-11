//WAP in go language to find the largest and smallest number in an array.

package main

import (
	"fmt"
)

func main() {
	var size, i, minPosition, maxPosition int

	fmt.Print("Enter the Array Size :")
	fmt.Scan(&size)

	arr := make([]int, size) //to make array as a dynamic content

	fmt.Printf("Enter the Array Items ::")
	for i = 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	largest := arr[0]
	smallest := arr[0]

	for i = 0; i < size; i++ {
		if largest < arr[i] {
			largest = arr[i]
			maxPosition = i
		}
		if smallest > arr[i] {
			smallest = arr[i]
			minPosition = i
		}
	}
	fmt.Println("\nThe Largest Number in this array = ", largest)
	fmt.Println("\nThe Index Position of Largest Number = ", maxPosition)
	fmt.Println("\nThe Smallest Number in this array = ", smallest)
	fmt.Println("\nThe Index Position of Smallest Number = ", minPosition)
}
