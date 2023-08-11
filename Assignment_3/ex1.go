package main

import (
	"fmt"
	"reflect"
)
func main(){
	
	var size int 
	fmt.Println("eneter the array size ")
	fmt.Scanln(&size)

	a := make([]int,size)
	fmt.Println(reflect.TypeOf(a))


	for i:=0;i<size;i++{
		fmt.Scanln(&a[i])
	}

	largest := a[0]
	Smallest := a[0]
	maxp :=0
	minp :=0

	for i := 0; i < size; i++ {
		if largest < a[i]{
			largest=a[i]
			maxp=i
		}

		if Smallest > a[i]{
			Smallest=a[i]
			minp=i
		}
	}

	fmt.Println("\nThe Largest Number in this array = ", largest)
	fmt.Println("\nThe Index Position of Largest Number = ", maxp)
	fmt.Println("\nThe Smallest Number in this array = ", Smallest)
	fmt.Println("\nThe Index Position of Smallest Number = ", minp)


}