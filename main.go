package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTicket int = 50

var conferenceName string = "Go Conference"
var remainingTicket int = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isValidEmail, isvalidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)

		if isValidName && isValidEmail && isvalidTicketNumber {

			bookTicket(userTicket, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The First names of bookings are: %v\n", firstNames)

			noTicketRemaining := remainingTicket == 0
			if noTicketRemaining {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name You Entered is Too Short")
			}
			if !isValidEmail {
				fmt.Println("Please Enter The Valid Email Address")
			}
			if !isvalidTicketNumber {
				fmt.Println("Number of Tickets You Entered Is Invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTicket, "tickets and", remainingTicket, "are still Remaining tickets")
	fmt.Println("Get your tickets here to attend the conference")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTicket int

	fmt.Println("Enter Your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your Email Address:")
	fmt.Scan(&email)

	fmt.Println("Enter Number Of Tickets:")
	fmt.Scan(&userTicket)
	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket int, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTicket
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("thank you %v %v for booking %v tickets you will recieved confirmation mail at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)

}
