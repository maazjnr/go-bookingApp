package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTicked int = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := greetUserInput()

	//user input validation function
	isValidEmail, isValidTicketNumber, isValidName := validUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		wg.Add(1)
		bookTicket(userTickets, firstName, lastName, email)

		go sendTicket(userTickets, firstName, lastName, email)

		//Calling the first userName Function
		firstNames := getFirstName()
		fmt.Printf("The first of the bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			//end the program if ticket is empty
			fmt.Println("Go conference ticket is booked out, come back next year")
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
	wg.Wait()
}

// Printing out the first name of the users who purchased the tickets
func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for user
	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are now remaining for %v\n", remainingTickets, conferenceName)
}

//greeting userInput function

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

func greetUsers() {
	fmt.Printf("Welcome to %v, booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",
		conferenceTicked, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// Sending the tickets to the user
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("send tickets:\n %v \nto email address %v\n", tickets, email)
	fmt.Println("#########################")
	wg.Done()
}
