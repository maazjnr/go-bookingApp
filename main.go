package main

import (
	"fmt"
)

func bookingApp() {

	conferenceName := "Go conference"
	const conferenceTicked int = 50
	var remainingTicket uint = 50

	fmt.Printf("Welcome to %v, booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",
		conferenceTicked, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	var bookings [50]string

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

	remainingTicket = remainingTicket - userTickets

	bookings[0] = firstName + " " + lastName

	fmt.Printf("Thank you %v %v for %v tickets. You will receive confirmation email at %v\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The type of the array: %T\n", bookings)
	fmt.Printf("The length of the array: %v\n", len(bookings))

	fmt.Printf("%v tickets are now remaining for %v\n", remainingTicket, conferenceName)
}

func main() {
	bookingApp()
}
