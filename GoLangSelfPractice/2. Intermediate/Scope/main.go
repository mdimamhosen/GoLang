package main

import (
	"firstgo/packageScope"
	"fmt"
)

var (
	x = 100
	z = 200
)

func main () {
	x := 10 // This will shadow the global variable x with the value 10 meaning that the global x will not be accessible in this scope anymore
	// <-- -->
	// := is used to declare and initialize a variable in the current scope
	// = is used to assign a value to a variable

	fmt.Println(x)
	{
		fmt.Println(x)
		y := 20
		fmt.Println(y)
	}
	// fmt.Println(y) // This will throw an error because y is not in scope

	// Accessing global variables
	fmt.Println(z)

	// Accessing package variables

	fmt.Println(packageScope.P)
}
