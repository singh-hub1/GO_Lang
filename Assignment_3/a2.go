//WAP in go language to accept the book details such as BookID, Title, Author, Price. Read and display the
// details of n number of books.

package main

import "fmt"

type Book_info struct {
	Book_id int
	Title   string
	Author  string
	Price   float32
}

func main() {
	var book [20]Book_info

	fmt.Println(" :: Enter how many book details you want:: ")
	var n int
	fmt.Scan(&n)

	fmt.Printf("\nEnter the Book information ::\n")
	for i := 0; i < n; i++ {
		fmt.Println("Enter the Book id ::")
		fmt.Scan(&book[i].Book_id)

		fmt.Println("Enter the Book title :: ")
		fmt.Scan(&book[i].Title)

		fmt.Println("Enter the Book Author :: ")
		fmt.Scan(&book[i].Author)

		fmt.Println("Enter the Book Price :: ")
		fmt.Scan(&book[i].Price)

		fmt.Println()
	}

	fmt.Printf("-------------------------------------------------------------------")
	fmt.Println("Display hthe book information:: ")
	for i := 0; i < n; i++ {
		fmt.Println("Book Id :: ", book[i].Book_id)
		fmt.Println("Book title :: ", book[i].Title)
		fmt.Println("Book Author :: ", book[i].Author)
		fmt.Println("Book Price :: ", book[i].Price)
		fmt.Println()
	}

}
