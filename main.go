package main

import "fmt"

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

	fmt.Println("Welcome to ", conferenceName, "booking application")
	// %v default format
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// infinite loop: interrupted by ctrl+c
	// allow users to repeat booking
	for {
		// Go is a statically typed language
		// You need to tell Go Compiler, the data type when declaring the variable.
		var userName string
		var userEmail string
		var userTickets uint

		// ask user for their name
		fmt.Print("Enter your name: ")
		fmt.Scan(&userName)
		// userName = "Tom"
		fmt.Print("Enter your email address: ")
		fmt.Scan(&userEmail)
		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		// userTickets = 2

		// remainingTickets = remainingTickets - uint(userTickets)
		remainingTickets -= userTickets
		bookings = append(bookings, userName)

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

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, userEmail)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings: %v\n", bookings)
	}

}
