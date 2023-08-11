//check wheather first string is substring of second string or not

package main

import (
	"fmt"
	s "strings"
)


func main(){
	var x string
	var y string
	var p=fmt.Println

	fmt.Println("Enter the first string:: ")
	fmt.Scanln(&x)
	
	fmt.Println("Enter the second string:: ")
	fmt.Scanln(&y)

	p("Contains:: ",s.Contains(x,y))
}                                       