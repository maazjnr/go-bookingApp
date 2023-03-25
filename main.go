package main

import (
	"fmt"
	"strings"
)

func bookingApp() {

	conferenceName := "Go conference"
	const conferenceTicked int = 50
	var remainingTickets uint = 50

	bookings := []string{}

	fmt.Printf("Welcome to %v, booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",
		conferenceTicked, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		isValidName := len(firstName) >= 0 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for %v tickets. You will receive confirmation email at %v\n",
				firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are now remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			if remainingTickets == 0 {
				//end the program if ticket is empty
				fmt.Println("Go conference ticket is booked out, come back next year")
				break
			}
		} else {
			fmt.Println("Your input data is incorrect, please try again")
		}

	}

}

func main() {
	bookingApp()
}
