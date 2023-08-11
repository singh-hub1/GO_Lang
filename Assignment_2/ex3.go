package  main

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

func main(){
	file,err:=os.Create("sudh.txt")
	  checkFileError(err)
	  defer file.Close()

	  length ,err := io.WriteString(file,"vishal singh nd aditi sharama")
	  checkFileError(err)
	  fmt.Printf("length of the file is %d",length)

	  readFile("sudh.txt")
}

func readFile(file string){
	dataByte,err := ioutil.ReadFile(file)
	checkFileError(err)
	fmt.Println(string(dataByte))
}


func checkFileError(err error){
if err!=nil{
	panic(err)
}
}