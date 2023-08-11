//3. Write a go program that creates a slice of integers, checks numbers from slice are even or odd and further sent to respective go routines through channel and display values received by go routines.

package main

import (
	"fmt"
	"time"
)

func main() {
	temp := make([]int, 5)
	even := make(chan int, 5) //buffered channel
	odd := make(chan int, 5)

	fmt.Println("Enter 5 Number : ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&temp[i])
	}

	fmt.Println("Slice : ", temp)

	go getOddEven(temp, even, odd)

	time.Sleep(time.Second)

	fmt.Println("Length Even : ", len(even))
	fmt.Println("Length Odd : ", len(odd))

	if len(even) != 0 {
		fmt.Println("Even : ")
		for i := 0; i <= len(even)+1; i++ {
			fmt.Print(<-even, " ")
		}

		fmt.Print(<-even, " ")

	}

	if len(odd) != 0 {
		fmt.Println("\nOdd : ")
		for i := 0; i <= len(odd)+1; i++ {
			fmt.Print(<-odd, " ")
		}
		fmt.Print(<-odd, " ")

	}
	time.Sleep(time.Second)

	fmt.Println()
}

func getOddEven(temp []int, even chan int, odd chan int) {

	for i := 0; i < 5; i++ {
		if temp[i]%2 == 0 {
			even <- temp[i]
		} else {
			odd <- temp[i]
		}

	}

}
