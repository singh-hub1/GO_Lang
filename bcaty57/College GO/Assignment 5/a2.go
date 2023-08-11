//2. WAP in GO program that executes 5 go routines simultaneously which generates numbers from 0 to 10, waiting between 0 and 250 ms after each go routine.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		go func() {
			rand.Seed(time.Now().Unix())
			fmt.Println("Number :", rand.Int31n(10))
		}()

		time.Sleep(time.Second)
	}
}
