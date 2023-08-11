//WAP in go language to create a file and write hello world in it and close the
// file by using defer statement.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("vishal.txt")

	checkNilErr(err)
	
	defer file.Close() //It will
	
	// fmt.Fprintln(file, "hello world")

	length, err := io.WriteString(file, "Success is my massive revenge!!")
	
	checkNilErr(err) //It will check an error

	fmt.Println("length is ", length)

	readFile("vishal.txt")

}

func readFile(file string) {
	dataByte, err := ioutil.ReadFile(file)
	checkNilErr(err)
	fmt.Println(string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
