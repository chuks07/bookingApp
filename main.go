package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)
var conferenceName = "Go conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}
	
	

func main() {




	 
	greetUsers()
	

	for {
		// ask user for 
		firstName,lastName, email, userTicket := getUserInput()



		
		isValidEmail, isValidname, isValidTicket := helper.ValidateUserInput( firstName, lastName,userTicket,remainingTickets,email)

		if isValidTicket && isValidEmail && isValidname {

			bookTickets(firstName,lastName,userTicket,email)		
			sendTicket(userTicket, firstName, lastName,email)

			firstNames := getFisrtName()
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

func greetUsers (){

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

}

func getFisrtName () []string {

	firstNames := []string{}
	// range defines the index and element of a slice or array
	for _, booking := range bookings {
	// we the fields action from the strings ackage to seperate words in a string
		
		
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
	
}

func getUserInput()(string, string, string, uint) {
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

		return firstName,lastName, email, userTicket
}

func bookTickets(firstName string, lastName string, userTicket uint, email string){

	remainingTickets= remainingTickets- userTicket

	var userData = UserData{
		firstName : firstName,
		lastName : lastName,
		email : email,
		numberOfTickets : userTicket, 
	}
	bookings= append(bookings, userData)
	
	fmt.Printf("The list of booking is %v\n", bookings)
		
	fmt.Printf("Thank you %v %v for buying %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("Thank you for your purchase there are/is %v left\n", remainingTickets)
	
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v  tickets were sent to %v %v\n", userTicket, firstName, lastName)

	fmt.Println("_________________________")
	fmt.Printf("sending tickets %v to %v\n", ticket, email)
	fmt.Println("_________________________")
}