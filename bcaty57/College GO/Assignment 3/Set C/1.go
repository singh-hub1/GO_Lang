//1. WAP in go language to accept two matrices and display it's multiplication.

package main

import "fmt"

func main() {
	var r1, r2, c1, c2 int

	fmt.Print("Enter Row of 1st matrix : ")
	fmt.Scan(&r1)

	fmt.Print("Enter Column of 1st matrix")
	fmt.Scan(&c1)

	fmt.Print("\nEnter Row of 2nd matrix : ")
	fmt.Scan(&r2)

	fmt.Print("Enter Column of 2nd matrix")
	fmt.Scan(&c2)

	if r1 != r2 || c1 != c2 {
		fmt.Println("\n Row and Column of both matrix must be same ..")
	}

	var m1 = make([][]int, r1)
	var m2 = make([][]int, r2)
	var result = make([][]int, r1)
	for i := 0; i < r1; i++ {
		m1[i] = make([]int, c1)
		m2[i] = make([]int, c2)
		result[i] = make([]int, c1)
	}

	fmt.Println("\nEnter 1st matrix : ")
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			fmt.Scan(&m1[i][j])
		}
	}

	fmt.Println("\nEnter 2nd matrix : ")
	for i := 0; i < r2; i++ {
		for j := 0; j < c2; j++ {
			fmt.Scan(&m2[i][j])
		}
	}

	fmt.Println("\nMultiplication of given Matrix : \n")
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			result[i][j] = 0

			for k := 0; k < r2; k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
			fmt.Print(result[i][j], " ")
		}
	}
}