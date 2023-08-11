//2. WAP in go language to print recursive sum of digits of given number.

package main

import "fmt"

func getSumOfDigit(num int) int {

	if num == 0 {
		return 0
	}
	return (num%10 + getSumOfDigit(int(num/10)))

}

func main() {
	var num, sum int
	sum = 0
	fmt.Print("Enter a no. : ")
	fmt.Scan(&num)

	sum = getSumOfDigit(num)

	fmt.Println("\nSum of Digit is : ", sum)
}