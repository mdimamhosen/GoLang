package main

import "fmt"

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

func main() {
	gameInitialization()
}
