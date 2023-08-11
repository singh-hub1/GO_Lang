//2. WAP in go language to accept the book details such as BookID, Title, Author, Price. Read and display the details of n number of books.

package main

import "fmt"

type Book struct {
	BookID        int
	Title, Author string
	Price         float32
}

func main() {
	var n int
	book := [100]Book{}

	fmt.Print("How many book details you wants to enter : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("Book ", i+1)
		fmt.Print("Book ID : ")
		fmt.Scan(&book[i].BookID)
		fmt.Print("Title : ")
		fmt.Scan(&book[i].Title)
		fmt.Print("Author : ")
		fmt.Scan(&book[i].Author)
		fmt.Print("Price : ")
		fmt.Scan(&book[i].Price)
		fmt.Println()
	}

	for i := 0; i < n; i++ {
		fmt.Println("Book ", i+1)
		fmt.Println("Book ID : ", book[i].BookID)
		fmt.Println("Title : ", book[i].Title)
		fmt.Println("Author : ", book[i].Author)
		fmt.Println("Price : ", book[i].Price)
		fmt.Println()
	}

}