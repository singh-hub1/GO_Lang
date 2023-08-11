//2. Write a program in go language to store n student information(rollno, name, percentage) and write a method to display student information in descending order of percentage.

package main

import (
	"fmt"
	"sort"
)

type Studnet struct {
	rollno     int
	name       string
	percentage float64
}

func main() {
	var n int

	fmt.Print("How many student details you wants to enter : ")
	fmt.Scan(&n)

	s := make([]Studnet, n)

	fmt.Println("\nEnter a Details : ")
	for i := 0; i < n; i++ {
		fmt.Println("\nStudent ", i+1)
		fmt.Print("Roll No. : ")
		fmt.Scan(&s[i].rollno)
		fmt.Print("Name : ")
		fmt.Scan(&s[i].name)
		fmt.Print("Percentage : ")
		fmt.Scan(&s[i].percentage)
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].percentage > s[j].percentage
	})

	// sort.Slice(s, func(i, j int) bool {
	// 	return s[i].rollno > s[j].rollno
	// })

	// sort.Slice(s, func(i, j int) bool {
	// 	return s[i].name > s[j].name
	// })

	fmt.Println("\nDecending sort by Student Percentage : ")

	for i := 0; i < n; i++ {
		fmt.Println("\nStudent ", i+1)
		fmt.Println("Roll No : ", s[i].rollno)
		fmt.Println("Name : ", s[i].name)
		fmt.Println("Percentage : ", s[i].percentage)
	}
}
