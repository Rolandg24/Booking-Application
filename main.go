package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 1250
	var remainingTickets = 1250

	fmt.Println("Welcome to the official", conferenceName, "Booking Application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets today to attend!")

}
