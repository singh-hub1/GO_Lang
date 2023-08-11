package main

import "fmt"

type shape interface{
	perimeter()float64
	area()float64
}

type rect struct{
	l,b float64
}

type cir  struct{
	radius float64
}

func( r rect) perimeter()(float64){
	return r.l*r.b
}

func( c cir) area()(float64){
	return c.radius
}

func main(){
 var r rect
 var c cir
var ch int32

 fmt.Print("1. Circle \n2. Rectangle\nEnter a Choice  : ")
	fmt.Scan(&ch)
	switch ch {
	case 1:
		fmt.Print("\nEnter a Radius : ")
		fmt.Scan(&c.radius)
		fmt.Println("\nArea of Circle : ", c.area())
		
	case 2:
		fmt.Print("\nEnter a length : ")
		fmt.Scan(&r.l)
		fmt.Print("Enter a breadth : ")
		fmt.Scan(&r.b)
		fmt.Println("\nPerimeter of Rectangle : ", r.perimeter())

}
}