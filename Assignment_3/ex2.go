package main

import "fmt"

func main(){
	// arr :=[]string{
	// 	"vishal",
	// 	"singh",
	// 	"arya",
	// }
	// fmt.Println(arr)

	// arr :=make([][]int,5)
	// for i := 0; i < 5; i++ {
	// 	arr[i] =make([]int,5)
	// }
	// fmt.Println(arr)

	arr:=make([]int,6)

	for i := 0; i < 6; i++ {
		fmt.Scanln(&arr[i])
	}


	for i := 1; i < 6; i++ {
		for j := 0; j < 6-i; j++ {
			if arr[j]>arr[j+1]{
				temp:=arr[j]
				arr[j]=arr[j+1]
				arr[j+1]=temp
			}
			
		}
	}

	fmt.Println("after shorting ::")
	for i := 0; i < 6; i++ {
		fmt.Println(arr[i])
	}
}