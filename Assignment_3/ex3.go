package main

import "fmt"

type Student struct{
	roll int
	name string 
	mark1,mark2,mark3,total,average float32
}
func main(){
var st[40] Student

for i := 0; i < 2; i++ {
	fmt.Println("\nStudent : ", i+1)
	fmt.Print(" Student Roll No. : ")
	fmt.Scan(&st[i].roll)

	fmt.Print("Student Name : ")
	fmt.Scan(&st[i].name)

	fmt.Print("Student Mark 1 : ")
	fmt.Scan(&st[i].mark1)

	fmt.Print("Student Mark 2 : ")
	fmt.Scan(&st[i].mark2)

	fmt.Print("Student Mark 3 : ")
	fmt.Scan(&st[i].mark3)

	st[i].total=st[i].mark1+st[i].mark2+st[i].mark3
	st[i].average=st[i].total/3


}


fmt.Println("\n Student Data are as follows :: ")
	for i := 0; i < 2; i++ {
		fmt.Println("Student : ", i+1)
		fmt.Println("Student Roll No. : ", st[i].roll)
		fmt.Println("Student Name : ", st[i].name)
		fmt.Println("Student Mark 1 : ", st[i].mark1)
		fmt.Println("Student Mark 2 : ", st[i].mark2)
		fmt.Println("Student Mark 3 : ", st[i].mark3)
		fmt.Println("Student Total marks : ", st[i].total)
		fmt.Println("Student Average marks  : ", st[i].average)
		fmt.Println()
	}

}