package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go conference"
	const conferenceTicked int = 50
	var remainingTickets uint = 50

	bookings := []string{}

	greetUsers(conferenceName, remainingTickets, conferenceTicked)

	for {

		firstName, lastName, email, userTickets := greetUserInput()

		//user input validation function
		isValidEmail, isValidTicketNumber, isValidName := validUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			//Calling the first userName Function
			firstNames := getFirstName(bookings)
			fmt.Printf("The first of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end the program if ticket is empty
				fmt.Println("Go conference ticket is booked out, come back next year")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("FirstName or LastName is too short")
			}

			if !isValidEmail {
				fmt.Println("Email is not valid because it is required @")
			}

			if !isValidTicketNumber {
				fmt.Println("Ticket number is not valid")
			}
		}

	}

}

//Printing out the first name of the users who purchased the tickets

func getFirstName(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}

func validUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (
	bool, bool, bool) {
	isValidName := len(firstName) >= 0 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidTicketNumber, isValidEmail
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for %v tickets. You will receive confirmation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are now remaining for %v\n", remainingTickets, conferenceName)
}

//greeting userInput fuction

func greetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your FirstName")
	fmt.Scan(&firstName)

	fmt.Println("Enter your LastName")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email")
	fmt.Scan(&email)

	fmt.Println("Enter number of Tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func greetUsers(conferenceName string, conferenceTicked uint, remainingTickets int) {
	fmt.Printf("Welcome to %v, booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",
		conferenceTicked, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
