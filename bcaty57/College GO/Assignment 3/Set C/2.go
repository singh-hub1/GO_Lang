// 2. WAP in go language to accept n records of employee information (eno,ename,salary) and display record of employees having maximum salary.

package main

import "fmt"

type Employee struct {
	eno    int
	ename  string
	salary float64
}

func main() {
	var n int
	var emp [10]Employee

	fmt.Print("How many employee details you wants to enter : ")
	fmt.Scan(&n)

	fmt.Println("Enter Details : ")

	for i := 0; i < n; i++ {
		fmt.Println("\nEmployee ", i+1)
		fmt.Print("No. : ")
		fmt.Scan(&emp[i].eno)
		fmt.Print("Name : ")
		fmt.Scan(&emp[i].ename)
		fmt.Print("Salary : ")
		fmt.Scan(&emp[i].salary)
	}

	loc := 0
	maxSalary := emp[0].salary

	for i := 1; i < n; i++ {
		if maxSalary < emp[i].salary {
			maxSalary = emp[i].salary
			loc = i
		}
	}

	fmt.Println("\nEmployee having maximum salary : \n")
	fmt.Println("No. : ", emp[loc].eno)
	fmt.Println("Name : ", emp[loc].ename)
	fmt.Println("Salary : ", emp[loc].salary)
}
