//1. WAP in go language to explain new function

package main

import "fmt"

func main() {
	var name, division, college_name string
	var rollno int

	fmt.Print("Enter Name :  ")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter Roll No. :  ")
	fmt.Scanf("%d", &rollno)
	fmt.Print("Enter division :  ")
	fmt.Scanf("%s", &division)
	fmt.Print("Enter College Name :  ")
	fmt.Scanf("%s", &college_name)

	fmt.Println("\nGiven Data is : ")
	fmt.Println("Name :  ", name)
	fmt.Println("Roll No.  :  ", rollno)
	fmt.Println("Division :  ", division)
	fmt.Println("College Name  :  ", college_name)

}
