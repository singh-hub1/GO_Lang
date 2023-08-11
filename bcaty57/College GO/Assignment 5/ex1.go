package main
import ("fmt"
"sync"
)

func main(){
	var num int
	fmt.Println("enter a number")
	fmt.Scan(&num)

	sqr:=make(chan int)
	cube:=make(chan int)

	 go calcSquares(num,sqr)
	 go calcCubes(num,cube)
}

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