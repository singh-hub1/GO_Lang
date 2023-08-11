//5. WAP in go language to explain new function

package main

import "fmt"

func main() {
	p := new(int)
	q := new(int)

	fmt.Print("Enter 1st no. : ")
	fmt.Scan(p)

	fmt.Print("Enter 2nd no. : ")
	fmt.Scan(q)

	add := *p + *q

	fmt.Println("\nAddition is : ", add)

}
