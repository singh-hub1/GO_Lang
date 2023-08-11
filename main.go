package main

import "fmt"

func main() {
	/*
		var name = "Go conference"
		const conferenceTicket = 50
		var remainningTicket = 50

		fmt.Printf("conferenceTicket is %T\n", conferenceTicket)
		fmt.Printf("remainningTicket is %T\n", remainningTicket)
		fmt.Printf("name is %T\n", name)

		fmt.Println("Welcome to", name, "booking application")
		fmt.Println("We have total of ", conferenceTicket, "ticket and ", remainningTicket, "are still available")

		fmt.Printf("welcome to %v booking application\n", name)
	*/

	////////////DATA TYPES//////////////////////////

	/*
		var userName string
		userName = "vishal singh"
		fmt.Println(userName)

		Name := "Singh Vishal"
		//This is another way to declare variable but not in the case of constant
		fmt.Print(Name)
	*/

	//////////Getting Input from User////////////////////////////////

	/*
		var firstName string
		var lastName string
		var emailId string
		var userTicket int

		fmt.Println("Enter your first name:: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email-id :: ")
		fmt.Scan(&emailId)

		fmt.Println("Enter your ticket you want:: ")
		fmt.Scan(&userTicket)

		fmt.Printf("Thank you %v %v for your  booking %v ticket.Confirmation will be sent to %v email", firstName, lastName, userTicket, emailId)

		remainningTicket = remainningTicket - userTicket
		fmt.Printf("\n%v ticket remainning for %v \n", remainningTicket, name)

	*/
	////////////////////ARRAY///////////////////////////

	//ARRAY is static in nature

	var bookings [50]string
	bookings[0] = "vishal singh"

	fmt.Printf("The whole array %v\n", bookings)
	fmt.Printf("The first value  %v\n", bookings[0])
	fmt.Printf("The array type  %T\n", bookings)
	fmt.Printf("The whole array length %v\n", len(bookings))

	////////////////SLICES///////////////////////////

	//slices is more dynamic in nature
	var booking []string
	NAME := "Vishal"
	booking = append(booking, NAME)

	fmt.Printf("The whole slice %v\n", booking)
	fmt.Printf("The first value  %v\n", bookings[0])
	fmt.Printf("The slices type  %T\n", booking)
	fmt.Printf("The whole slice length %v\n", len(booking))

	////////////////LOOPS ///////////////////////////////
	for i := 1; i <= 100; i++ {
		fmt.Print(i, " ")
	}

}
