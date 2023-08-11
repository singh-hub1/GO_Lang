//2. WAP in go language to create a file and write hello world in it and close the file by using defer statement.

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")

	defer file.Close()

	if err != nil {
		fmt.Println("\nERROR ....")
	} else {
		fmt.Fprintln(file, "Hello World !!!")

	}

}
