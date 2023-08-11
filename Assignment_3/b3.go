//3. WAP in go language to accept n student details like roll_no, stud_name, mark1,mark2, mark3. Calculate the total and average of marks using structure.

package main

import "fmt"

type Student struct {
	roll_no   int
	stud_name  string
	mark1, mark2, mark3, total, avg float64
}

func main() {
	var n int
	var st [10]Student

	fmt.Print("How many Student details you wants to enter :")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("\nStudent : ", i+1)
		fmt.Print(" Student Roll No. : ")
		fmt.Scan(&st[i].roll_no)

		fmt.Print("Student Name : ")
		fmt.Scan(&st[i].stud_name)

		fmt.Print("Student Mark 1 : ")
		fmt.Scan(&st[i].mark1)

		fmt.Print("Student Mark 2 : ")
		fmt.Scan(&st[i].mark2)

		fmt.Print("Student Mark 3 : ")
		fmt.Scan(&st[i].mark3)

		st[i].total = st[i].mark1 + st[i].mark2 + st[i].mark3
		st[i].avg = st[i].total / 3
	}

	fmt.Println("\n Student Data are as follows :: ")
	for i := 0; i < n; i++ {
		fmt.Println("Student : ", i+1)
		fmt.Println("Student Roll No. : ", st[i].roll_no)
		fmt.Println("Student Name : ", st[i].stud_name)
		fmt.Println("Student Mark 1 : ", st[i].mark1)
		fmt.Println("Student Mark 2 : ", st[i].mark2)
		fmt.Println("Student Mark 3 : ", st[i].mark3)
		fmt.Println("Student Total marks : ", st[i].total)
		fmt.Println("Student Average marks  : ", st[i].avg)
		fmt.Println()
	}

}
