package main

import (
	"fmt"
	"strings"
	"time"
)



func main() {

	// Calculator UI
	for {
		fmt.Println(strings.Repeat("=", 30))
		fmt.Println("          Console Calculator")
		fmt.Println(strings.Repeat("=", 30))
		fmt.Println("1. Addition (+)")
		fmt.Println("2. Subtraction (-)")
		fmt.Println("3. Multiplication (*)")
		fmt.Println("4. Division (/)")
		fmt.Println("5. Exit")
		fmt.Println(strings.Repeat("=", 30))

		// Read user choice
		var choice int
		fmt.Print("Choose an operation (1-5): ")
		fmt.Scan(&choice)

		if choice == 5 {
			fmt.Println("Exiting calculator. Goodbye!")
			time.Sleep(1 * time.Second)
			break
		}

		var num1, num2 float64
		fmt.Print("Enter first number: ")
		fmt.Scan(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)

		// Perform operation

		fmt.Println(strings.Repeat("=", 30))
		fmt.Println("          Console Calculator")
		fmt.Println(strings.Repeat("=", 30))

		switch choice {
		case 1:
			fmt.Printf("Result: %f + %f = %f\n", num1, num2, num1+num2)
		case 2:
			fmt.Printf("Result: %f - %f = %f\n", num1, num2, num1-num2)
		case 3:
			fmt.Printf("Result: %f * %f = %f\n", num1, num2, num1*num2)
		case 4:
			if num2 == 0 {
				fmt.Println("Error: Cannot divide by zero!")
			} else {
				fmt.Printf("Result: %f / %f = %f\n", num1, num2, num1/num2)
			}
		default:
			fmt.Println("Invalid choice! Please select a valid operation.")
		}

		fmt.Println(strings.Repeat("=", 30))
		fmt.Print("Press Enter to continue...")
		fmt.Scanln()
		fmt.Scanln()

	}
}
