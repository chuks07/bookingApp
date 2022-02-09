package main

import (
	"fmt"
	"strings"
)

	
	

func main() {

	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings [] string 
	 
	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	

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

		isValidEmail, isValidname, isValidTicket := validateUserInput( firstName, lastName,userTicket,remainingTickets,email)

		if isValidTicket && isValidEmail && isValidname {
			remainingTickets= remainingTickets- userTicket
			bookings= append(bookings, firstName + " "+ lastName)
		
		
		
			fmt.Printf("Thank you %v %v for buying %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
			fmt.Printf("Thank you for your purchase there are/is %v left\n", remainingTickets)

		firstNames := getFisrtName(bookings)
		fmt.Printf("These are all the bookings: %v\n", firstNames)


			noTicketRemaining := remainingTickets == 0
			if noTicketRemaining {
				fmt.Printf("The %v is completely booked up\n", conferenceName)
				break
			}			
		} else 	{
			if !isValidname{
				fmt.Println("The name is too short")
			}
			if !isValidEmail{
				fmt.Println("Email input is invalid")
			}
			if !isValidTicket{
				fmt.Println("Ticket input is invalid")
			}
	
		}
	}
}

func greetUsers (conferenceName string, conferenceTickets int, remainingTickets uint){

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

}

func getFisrtName (bookings []string) []string {

	firstNames := [] string{}
	// range defines the index and element of a slice or array
	for _, booking := range bookings {
	// we the fields action from the strings ackage to seperate words in a string
		var name = strings.Fields(booking)
		var firstName = name[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
	
}

func validateUserInput (firstName string, lastName string, userTicket uint, remainingTickets uint, email string ) (bool, bool ,bool ) {

	isValidname := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email,"@")
		isValidTicket := userTicket > 0 && userTicket <= remainingTickets

		return isValidEmail, isValidname, isValidTicket
}