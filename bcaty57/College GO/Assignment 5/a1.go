package main

import "fmt"

func calcSquares(number int, square chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	square <- sum
}

func calcCubes(number int, cube chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cube <- sum
}

func main() {
	var number int
	sqr := make(chan int)
	cube := make(chan int)

	fmt.Println("Enter a Integer : ")
	fmt.Scan(&number)

	go calcSquares(number, sqr)
	go calcCubes(number, cube)
	squares, cubes := <-sqr, <-cube
	fmt.Println("\nSum of Square : ", squares)
	fmt.Println("Sum of Cubes : ", cubes)
	fmt.Println("Final output", squares+cubes)
}
