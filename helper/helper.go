package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, userTicket uint, remainingTickets uint, email string) (bool, bool, bool) {

	isValidname := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTicket > 0 && userTicket <= remainingTickets

	return isValidEmail, isValidname, isValidTicket
}