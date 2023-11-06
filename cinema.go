package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)


var users = make(map[string]User)
var factory = MovieFactory{}
var movies = []Movie{

}
var currentUser User

func main() {

	fmt.Println("Welcome to the Cinema Ticket Booking System")

	for {
		fmt.Println("\n1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Print("Please select an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			login()
		case 2:
			register()
		case 3:
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
