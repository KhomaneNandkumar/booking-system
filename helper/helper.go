package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTicket int, remainingTicket int) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isvalidTicketNumber := userTicket > 0 && userTicket <= remainingTicket
	return isValidName, isValidEmail, isvalidTicketNumber
}
