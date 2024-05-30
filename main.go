package main

import "fmt"

func main() {

	conferenceName := "8-Bit Gaming Conference"
	const conferenceTickets int = 1250
	var remainingTickets int = 1250

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to the official %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets today to attend!\n")

	var userName string
	var userTickets int
	//asking user for their name here.
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("%v booked %v tickets. \n", userName, userTickets)

}
