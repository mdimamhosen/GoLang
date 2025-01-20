package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("Choose a difficulty level (1 to 3):")
	fmt.Println("1. Easy (1-10)")
	fmt.Println("2. Medium (1-50)")
	fmt.Println("3. Hard (1-100)")

	var choice int
	fmt.Scan(&choice)

	var maxNumber int
	switch choice {
	case 1:
		maxNumber = 10
	case 2:
		maxNumber = 50
	case 3:
		maxNumber = 100
	default:
		fmt.Println("Invalid choice! Defaulting to Easy level.")
		maxNumber = 10
	}

	var computerRand int
	rand.Seed(time.Now().UnixNano())
	computerRand = rand.Intn(maxNumber) ;

	var userGuess int
	var attempts int
	fmt.Print("Number of times you want to play: ")
	fmt.Scan(&attempts)
	for(attempts > 0){
		fmt.Printf("Guess the number (1-%d): ", maxNumber)
		fmt.Scan(&userGuess)

		if userGuess < computerRand {
			fmt.Println("Too low!")
		} else if userGuess > computerRand {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Congratulations! You got it!")
			break
		}
		attempts--
	}


}
