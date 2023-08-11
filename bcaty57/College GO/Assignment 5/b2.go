//Wap in Go main go routine to read and write fibonacci series to the channel.


package main

import (
	"fmt"
)

func fibonacci(ch chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x: // write to channel ch
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 10

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch) // read from channel ch
		}
		quit <- false
	}(n)

	fibonacci(ch, quit)
}