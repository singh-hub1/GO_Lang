//WAP in go language to Initialize a Slice using Multi-Line Syntax and display

package main

import (
	"fmt"
	// "reflect"
)

func main() {
	//var multi[]float64
	//fmt.Println(reflect.ValueOf(multi).Kind())

	// var mul=make([]int,10)//here length and capacity is same
	// var mul_1=make([]int,10,15)//here length is 10 and capacity is 15

	// fmt.Printf("len::%v\t\t cap::%v\n",len(mul_1),cap(mul_1));
	// fmt.Printf("len::%v\t\t cap::%v\n",len(mul),cap(mul));

	///////////////////Slice using some slice literals////////////////
	// var mul = []int{10, 20, 30, 40, 50}
	// fmt.Printf("length ::%v\t\t Capacity ::%v\n", len(mul), cap(mul))

	/////////declare slice using new keyboard///////////////

	// var mul = new([50]int)[0:10]
	// fmt.Println(reflect.ValueOf(mul).Kind())
	// fmt.Printf("length ::%v\t\t Capacity ::%v\n", len(mul), cap(mul))
	// fmt.Print(mul)

	//////to insert some elements in the end of the array/////

	// a := make([]int, 2, 5)
	// a[0] = 10
	// a[1] = 20
	// a = append(a, 30, 40, 50, 60, 70)
	// fmt.Printf("length ::%v\t\t Capacity ::%v\n", len(a), cap(a))
	// fmt.Print(a)

	//INitialize a slice using multiline syntax////////////////

	names := []string{
		"vishal",
		"aditya",
		"justin",
	}

	fmt.Println(names)

}
