package helper

import "strings"

// export function explicity -> capitalize the first letter
func ValidateUserInput(firstName string, lastName string, userEmail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// validate user input
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // positive number + greater than 0
	return isValidName, isValidEmail, isValidTicketNumber
}
