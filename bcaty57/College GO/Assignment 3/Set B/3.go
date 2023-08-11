//3. WAP in go language to accept n student details like roll_no, stud_name, mark1,mark2, mark3. Calculate the total and average of marks using structure.

package main

import "fmt"

type Student struct {
	roll_no                         int
	stud_name                       string
	mark1, mark2, mark3, total, avg float32
}

func main() {
	var n int
	var student [10]Student

	fmt.Print("How many Student details you wants to enter :")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("\nStudent : ", i+1)
		fmt.Print("Roll No. : ")
		fmt.Scan(&student[i].roll_no)
		fmt.Print("Name : ")
		fmt.Scan(&student[i].stud_name)
		fmt.Print("Mark 1 : ")
		fmt.Scan(&student[i].mark1)
		fmt.Print("Mark 2 : ")
		fmt.Scan(&student[i].mark2)
		fmt.Print("Mark 3 : ")
		fmt.Scan(&student[i].mark3)
		student[i].total = student[i].mark1 + student[i].mark2 + student[i].mark3
		student[i].avg = student[i].total / 3
	}

	fmt.Println("\nGiven data is : ")
	for i := 0; i < n; i++ {
		fmt.Println("Student : ", i+1)
		fmt.Println("Roll No. : ", student[i].roll_no)
		fmt.Println("Name : ", student[i].stud_name)
		fmt.Println("Mark 1 : ", student[i].mark1)
		fmt.Println("Mark 2 : ", student[i].mark2)
		fmt.Println("Mark 3 : ", student[i].mark3)
		fmt.Println("Total : ", student[i].total)
		fmt.Println("Average : ", student[i].avg)
		fmt.Println()
	}

}