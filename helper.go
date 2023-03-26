package main

import (
	"strings"
)

func validUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (
	bool, bool, bool) {
	isValidName := len(firstName) >= 0 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidTicketNumber, isValidEmail
}
