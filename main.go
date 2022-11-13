package main

import (
	"fmt"
	"strings"
)

func main() {
	// Type Inference: Go can infer the type when you assign a value
	conferenceName := "Go Conference" // syntactic suger
	const conferenceTickets = 50      // total of attendee
	var remainingTickets uint = 50    // non-negative
	// var variable_name [size]variable_type
	// array: fixed size
	// var bookings [50]string // attendees
	// slice (abstraction of array): dynamic size
	bookings := []string{}

	// %T prints data type
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// infinite loop: interrupted by ctrl+c
	// allow users to repeat booking
	for {
		// Go is a statically typed language
		// You need to tell Go Compiler, the data type when declaring the variable.
		var firstName string
		var lastName string
		var userEmail string
		var userTickets uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&userEmail)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// validate user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // positive number + greater than 0

		if isValidName && isValidEmail && isValidTicketNumber {
			// remainingTickets = remainingTickets - uint(userTickets)
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Println(remainingTickets)  // print value of remainingTickets
			// fmt.Println(&remainingTickets) // print memory address of remainingTickets

			// fmt.Printf("The whole array: %v\n", bookings)   // [Nick ]
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Array type: %T\n", bookings)        // [50]string
			// fmt.Printf("Array length: %d\n", len(bookings)) // 50
			// fmt.Printf("The whole slice: %v\n", bookings)   // [Nick]
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)        // []string
			// fmt.Printf("Slice length: %d\n", len(bookings)) // 1

			fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, userTickets, userEmail)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// call function print first names
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// boolean
			var noTicketsRemaning bool = remainingTickets == 0
			if noTicketsRemaning {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			// give user a useful feedback
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}

}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Println("Welcome to ", confName, "booking application")
	// %v default format
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	// for-each
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
