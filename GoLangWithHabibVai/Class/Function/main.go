package main

import "fmt"

// add1 takes two integers, calculates their sum, and prints it.
func add1(x int, y int) {
	sum := x + y
	fmt.Println(sum)
}

// add2 takes two integers and returns their sum.
func add2(x int, y int) int {
	return x + y
}

// add_mul takes two integers and returns their sum and product.
func add_mul(x int, y int) (int, int) {
	sum := x + y
	mul := x * y
	return sum, mul
}

// printSomething prints a predefined message.
func printSomething() {
	fmt.Println("Something is printed")
}

// sayHello takes a string and prints a greeting message.
func sayHello(str string) {
	fmt.Println("Hello,", str)
}

// subtract takes two integers and returns their difference.
func subtract(x int, y int) int {
	return x - y
}

// divide takes two integers and returns their quotient.
func divide(x int, y int) float64 {
	if y != 0 {
		return float64(x) / float64(y)
	}
	fmt.Println("Error: Division by zero")
	return 0
}

func main() {
	// Call add1 with arguments 42 and 13, and print the sum.
	add1(42, 13)

	// Print the result of add2(42, 13) to the console.
	fmt.Println(add2(42, 13))

	// Call add_mul with arguments 42 and 13, and print the sum and product.
	sum, mul := add_mul(42, 13)
	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", mul)

	// Call printSomething to print a predefined message.
	printSomething()

	// Call sayHello with the argument "World" to print a greeting message.
	sayHello("World")

	// Call subtract with arguments 42 and 13, and print the difference.
	fmt.Println("Difference:", subtract(42, 13))

	// Call divide with arguments 42 and 13, and print the quotient.
	fmt.Println("Quotient:", divide(42, 13))
}

// This are standard functions that are used in the main function.
