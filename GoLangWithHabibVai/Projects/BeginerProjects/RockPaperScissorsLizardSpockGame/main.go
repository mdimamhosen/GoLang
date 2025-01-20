package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func gameInitialization() {
	fmt.Println("Welcome to Rock Paper Scissors Lizard Spock Game!")
	fmt.Print("Please enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Print("Hello ", name, ", let's start the game! ")

	// Explaining the rules in two lines
	fmt.Println("\nRules of the Game: Rock crushes Scissors and Lizard, Scissors cuts Paper and decapitates Lizard, Paper covers Rock and disproves Spock, Lizard eats Paper and poisons Spock, Spock vaporizes Rock and smashes Scissors.")
	fmt.Println("Please choose one of the following options: Rock, Paper, Scissors, Lizard, Spock")
}

func baseUiInitialization() {
	fmt.Println("\nGame Options:")
	fmt.Println("   Rock       Paper      Scissors       Lizard       Spock")
	fmt.Println(strings.Repeat("-", 60))
}

func makeComputerChoice() string {
	options := []string{"Rock", "Paper", "Scissors", "Lizard", "Spock"}
	computerChoice := options[rand.Intn(len(options))]
	return computerChoice
}
func optionSelection() {
	options := []string{"Rock", "Paper", "Scissors", "Lizard", "Spock"}
	fmt.Print("Enter your choice (name or number): ")
	var input string
	fmt.Scanln(&input)

	// Determine if input is a number or string
	var choice string
	// Try to parse the input as a number
	var num int
	if _, err := fmt.Sscanf(input, "%d", &num); err == nil && num > 0 && num <= len(options) {
		// If it's a valid number, select the corresponding option
		choice = options[num-1]
	} else {
		// Otherwise, handle it as a string input
		// Normalize input by converting it to lowercase
		choice = strings.Title(strings.ToLower(input)) // Capitalize first letter, make the rest lowercase
	}

	// Check if the choice is valid
	if contains(options, choice) {
		fmt.Println("You selected:", choice)
	} else {
		fmt.Println("Invalid choice. Please try again.")
		optionSelection() // Retry the selection if invalid
		return
	}

	// Make computer choice and check winner
	computerChoice := makeComputerChoice()
	checkWinner(choice, computerChoice)
}


// Helper function to check if a slice contains a value
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func checkWinner(playerChoice string, computerChoice string) {
	fmt.Println("Computer selected:", computerChoice)
	if playerChoice == computerChoice {
		fmt.Println("It's a tie!")
	} else {
		// Determine the winner based on the rules
		switch playerChoice {
		case "Rock":
			if computerChoice == "Scissors" || computerChoice == "Lizard" {
				fmt.Println("You win!")
			} else {
				fmt.Println("Computer wins!")
			}
		case "Paper":
			if computerChoice == "Rock" || computerChoice == "Spock" {
				fmt.Println("You win!")
			} else {
				fmt.Println("Computer wins!")
			}
		case "Scissors":
			if computerChoice == "Paper" || computerChoice == "Lizard" {
				fmt.Println("You win!")
			} else {
				fmt.Println("Computer wins!")
			}
		case "Lizard":
			if computerChoice == "Spock" || computerChoice == "Paper" {
				fmt.Println("You win!")
			} else {
				fmt.Println("Computer wins!")
			}
		case "Spock":
			if computerChoice == "Rock" || computerChoice == "Scissors" {
				fmt.Println("You win!")
			} else {
				fmt.Println("Computer wins!")
			}
		}
	}
}

func main() {
	gameInitialization()
	baseUiInitialization()
	optionSelection()
}
