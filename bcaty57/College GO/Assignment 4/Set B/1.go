//1. Write a program in go language to create structure student. Write a method show() whose receiver is a pointer of struct student.

package main

import "fmt"

type student struct {
	roll  int
	name  string
	marks float64
}

func show(obj *student) {
	fmt.Println("\nRoll No. : ", obj.roll)
	fmt.Println("Name : ", obj.name)
	fmt.Println("Marks : ", obj.marks)
}
func main() {
	var s student

	fmt.Print("Student Roll No. : ")
	fmt.Scan(&s.roll)
	fmt.Print("Student Name : ")
	fmt.Scan(&s.name)
	fmt.Print("Student Marks : ")
	fmt.Scan(&s.marks)

	show(&s)
	
}
