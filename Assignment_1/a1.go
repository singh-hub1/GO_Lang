//WAP in go language to print Student name, rollno, division and college name
package main

import (
	"fmt"
)

func main() {
	var fname string
	var lname string
	var rollNo int
	var division string
	var clgName string

	fmt.Println("Enter your first name :: ")
	fmt.Scanln(&fname)

	fmt.Println("Enter your last name :: ")
	fmt.Scanln(&lname)

	fmt.Println("Enter the roll number :: ")
	fmt.Scanln(&rollNo)

	fmt.Println("Enter the division :: ")
	fmt.Scanln(&division)

	fmt.Println("Enter the college name :: ")
	fmt.Scanln(&clgName)

	fmt.Printf("Welcome  %v %v your roll number is %v with your division is %v in the %v", fname, lname, rollNo, division, clgName)
}
