//check wheather given number is single digit or not 

package main

import "fmt"

func main(){
	var a int 

	fmt.Println("Enter any number:: ")
	fmt.Scanln(&a)

	if (a >= 0 && a<=9) {
		fmt.Println("given number is single didit:: ")
	}else if(a>=10 && a<=99){
		fmt.Println("given number is not single digit::")
	}
}