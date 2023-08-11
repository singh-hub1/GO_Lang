package main

import(
	"fmt"
	"strconv"
)
func findType(a interface{}){
	switch a.(type){
		case string:fmt.Println("\n It is a String. ")
	case int:
		fmt.Println("\nIt is a Integer. ")
	case float64:
		fmt.Println("\nIt is a Float ")
	default:
		fmt.Println("Unknown Data Type ..")
	}
}

func main(){
	var ch interface{}
	var temp string

	fmt.Println("enter anything")
	fmt.Scan(&temp)

	if v,err := strconv.Atoi(temp)
		err==nil{
			ch=v
		}else if v,err:=strconv.ParseFloat(temp,32) //32 bit of float 32 
		  err==nil{
			  ch=v
		  }else{
			  ch=v
		  }
		  findType(ch)

}