//3. WAP in go language to demonstrate working of slices (like append, remove, copy etc.)

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := make([]int, len(arr))

	fmt.Println("Slice : ", arr)

	arr = append(arr, 10)
	fmt.Println("\nAfter appending 10 in slice : ", arr)

	// arr = append(arr[:len(arr)-1], arr[:len(arr)-2]...)
	arr = arr[:len(arr)-1]
	fmt.Println("\nAfter removing last element in slice : ", arr)

	copy(arr2, arr)
	fmt.Println("\nCopying one slice into another one :", arr2)

	// arr.remove(2)
	fmt.Println("removing this element from array arr:: ",arr);
}
