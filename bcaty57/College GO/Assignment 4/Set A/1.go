//1. Write a program in go language to create an interface shape that includes area and perimeter. Implements these methods in circle and rectangle type.

package main

import (
	"fmt"
)

type Area interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	length, breadth float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	var r Rectangle
	var c Circle
	var ch int

	fmt.Print("1. Circle \n2. Rectangle\nEnter a Choice  : ")
	fmt.Scan(&ch)
	switch ch {
	case 1:
		fmt.Print("\nEnter a Radius : ")
		fmt.Scan(&c.radius)
		fmt.Println("\nArea of Circle : ", c.area())
		fmt.Println("\nPerimeter  of Circle : ", c.perimeter())
	case 2:
		fmt.Print("\nEnter a length : ")
		fmt.Scan(&r.length)
		fmt.Print("Enter a breadth : ")
		fmt.Scan(&r.breadth)
		fmt.Println("\nArea of Rectangle : ", r.area())
		fmt.Println("\nPerimeter of Rectangle : ", r.perimeter())

	}
}
