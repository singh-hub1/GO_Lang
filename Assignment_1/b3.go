//WAP in go language to print Fibonacci series of n terms.
package main

import (
	"fmt"
)

func main() {
	var fib int
	fmt.Println("Enter the terms you want:: ")
	fmt.Scanln(&fib)
	Fibonacciseries(fib)

}

func Fibonacciseries(fib int) {
	var n3,n1,n2 int = 0,0,1

	for i:= 1;i<=fib;i++{
		
		fmt.Println(n1)
		n3=n1+n2
		n1=n2
		n2=n3
	}
}
