//WAP in go to print table of given number

package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter any number for table:: ")
	fmt.Scanln(&a)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", a, i, a*i)
	}

}
