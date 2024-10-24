package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type User struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var bookings = make([]User, 0)

var wg = sync.WaitGroup{}

func main() {

	conferenceName := "Go"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	var email string
	var firstName string
	var lastName string
	var userTickets uint

	greetUsers(conferenceName, remainingTickets)

	for {

		getData(&firstName, &lastName, &email, &userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2

		isValidEmail := strings.Contains(email, "@")

		if !isValidName || !isValidEmail {
			fmt.Println("Incorrect imput type")
			continue
		}

		if userTickets > remainingTickets {
			fmt.Printf("Whe only got more %v tickets, please buy a lower amount\n", remainingTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets

		var userData = User{
			firstName: firstName,
			lastName:  lastName,
			email:     email,
			tickets:   userTickets,
		}

		bookings = append(bookings, userData)
		printFirstNames()
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		printFinalData(firstName, lastName, remainingTickets, userTickets)
		fmt.Printf("Info of the guests: %v\n", bookings)

		if !checkRemainingTickets(remainingTickets) {
			fmt.Println("The tickets have sold out, thank you for the desire")
			break
		}

	}
	wg.Wait()
}

func greetUsers(conferenceName string, remainTickets uint) {
	fmt.Println("Welcome to our", conferenceName, "conference")
	fmt.Printf("We have a total of %v remaining tickets\n", remainTickets)
	fmt.Println("Take your ticket at the end of the corridor")
}
func getData(firstName *string, lastName *string, email *string, ticketQuantity *uint) {
	fmt.Print("First Name: ")
	fmt.Scan(firstName)

	fmt.Print("Last Name: ")
	fmt.Scan(lastName)

	fmt.Print("Email: ")
	fmt.Scan(email)

	fmt.Print("Tickets quantity: ")
	fmt.Scan(ticketQuantity)
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("The names and of guests are %v\n", firstNames)
}

func checkRemainingTickets(remainingTickets uint) bool {
	haveRemainingTickets := remainingTickets != 0
	return haveRemainingTickets
}

func printFinalData(firstName string, lastName string, remainingTickets uint, userTickets uint) {
	fmt.Printf("User %v %v booked %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#############")
	fmt.Println("Sending ticket:", ticket, "\nto email address:", email)
	fmt.Println("#############")
	wg.Done()
}
