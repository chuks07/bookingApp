package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings [] string 


	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for {
		// ask user for input

		var firstName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Println("Enter your firstname:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastname:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("How any ticket(s) do you want:")
		fmt.Scan(&userTicket)

		remainingTickets= remainingTickets- userTicket
		bookings= append(bookings, firstName + " "+ lastName)
		
		
		
		fmt.Printf("Thank you %v %v for buying %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
		fmt.Printf("Thank you for your purchase there are/is %v left\n", remainingTickets)
		fmt.Printf("These are all the bookings: %v", bookings)

	}
}
