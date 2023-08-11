//3. WAP in go language using function to check whether accepts number is palindrome or not.

package main

import "fmt"

func palindrome(num int) (reverse int) {

	for num != 0 {
		rem := num % 10
		reverse = reverse*10 + rem
		num /= 10
	}
	return
}

func main() {
	var number int
	var reverse int = 0

	fmt.Print("Enter any positive integer : ")
	fmt.Scan(&number)

	reverse = palindrome(number)

	if number == reverse {
		fmt.Println("\n", number, "is a Palindrome")
	} else {
		fmt.Println("\n", number, "is not a Palindrome")
	}

}
