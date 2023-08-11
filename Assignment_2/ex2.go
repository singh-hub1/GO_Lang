package main

import "fmt"

func main(){

	var a int 
	fmt.Println("Enetr the number")
	fmt.Scanln(&a)

	var x,y=mul(a)
	fmt.Println("the addition is ",x,"as")
	fmt.Println("the subtraction is ",y,"as")
}

func mul(a int )(add int,sub int){
    add = a+2
	sub = a-2
	return 
}