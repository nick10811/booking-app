package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

// package level variable (global)
// Type Inference: Go can infer the type when you assign a value
var conferenceName = "Go Conference" // syntactic suger
const conferenceTickets = 50         // total of attendee
var remainingTickets uint = 50       // non-negative
// var variable_name [size]variable_type
// array: fixed size
// var bookings [50]string // attendees
// slice (abstraction of array): dynamic size
// var bookings = []string{}
// var bookings = make([]map[string]string, 0) // list of maps
var bookings = make([]UserData, 0)

// mixed data type by structure
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	// %T prints data type
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	greetUsers()

	// infinite loop: interrupted by ctrl+c
	// allow users to repeat booking
	for {
		firstName, lastName, userEmail, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, userEmail)
			// make a concurrent process
			go sendTicket(userTickets, firstName, lastName, userEmail)

			// call function print first names
			firstNames := getFirstNames()
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

func greetUsers() {
	fmt.Println("Welcome to ", conferenceName, "booking application")
	// %v default format
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	// for-each
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, userEmail, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	// remainingTickets = remainingTickets - uint(userTickets)
	remainingTickets -= userTickets

	// create a map for a user
	// map[key]value
	// Go can't mix data type in map
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = userEmail
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           userEmail,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

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
}

// simulate to send ticket
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("##########")
}
