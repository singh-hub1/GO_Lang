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

	fmt.Println("Enter Employee  Details : ")

	for i := 0; i < n; i++ {
		fmt.Println("\nEmployee ", i+1)

		fmt.Print("Enter Employee No. : ")
		fmt.Scan(&emp[i].eno)

		fmt.Print("Enter Employee Name : ")
		fmt.Scan(&emp[i].ename)

		fmt.Print("Enter Employee Salary : ")
		fmt.Scan(&emp[i].salary)

		fmt.Println()
	}
	loc:=0
	max :=emp[0].salary

	for i := 1; i < n; i++ {

		if max < emp[i].salary {
			max=emp[i].salary
			loc=i
		}
		
	}

	fmt.Println("\nEmployee having maximum salary :")
	fmt.Println("Employee No. : ", emp[loc].eno)
	fmt.Println("Employee Name : ", emp[loc].ename)
	fmt.Println("Employee Salary : ", emp[loc].salary)
}
