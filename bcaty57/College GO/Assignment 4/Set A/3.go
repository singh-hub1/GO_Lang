//3. Write a program in go language to create structure author. Write a method show() whose receiver is struct author.

package main

import "fmt"

type author struct {
	name     string
	bookName string
}

func (obj author)show() {
	fmt.Println("\nAuthor Name : ", obj.name)
	fmt.Println("Book Name : ", obj.bookName)
}

func main() {
	var a author

	fmt.Print("Enter a Author name : ")
	fmt.Scan(&a.name)
	fmt.Print("Enter a Book Name : ")
	fmt.Scan(&a.bookName)

	a.show()

}
