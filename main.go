package main

import "fmt"

func main() {
	// Type Inference: Go can infer the type when you assign a value
	conferenceName := "Go Conference" // syntactic suger
	const conferenceTickets = 50      // total of attendee
	var remainingTickets uint = 50    // non-negative

	// %T prints data type
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to ", conferenceName, "booking application")
	// %v default format
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Go is a statically typed language
	// You need to tell Go Compiler, the data type when declaring the variable.
	var userName string
	var userEmail string
	var userTickets int

	// ask user for their name
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)
	// userName = "Tom"
	fmt.Print("Enter your email address: ")
	fmt.Scan(&userEmail)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	// userTickets = 2

	// fmt.Println(remainingTickets)  // print value of remainingTickets
	// fmt.Println(&remainingTickets) // print memory address of remainingTickets

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, userEmail)

}
