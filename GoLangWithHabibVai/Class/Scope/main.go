// <!-- Scope in Go Lang -->
package main

import "fmt"

// let us declare a package level variable a and a function this is also called global scope
var a int = 100

func printSomething() {
	// let us declare a function level variable b this is also called local scope
	var b int = 200
	fmt.Println(b) // 200

	fmt.Println(a) // 100

}

func addSomething(a int , b int ) int {
	// let us declare a function level variable c this is also called local scope
	var c int = a + b
	return c
}

func main() {

	// let us declare a variable
	var x int = 10
	var y int = 20

	// let us declare a block
	{
		// let us declare a variable
		var x int = 30
		var y int = 40
		var z int = 50

		fmt.Println(x, y, z) // 30 40 50
	}
	fmt.Println(x, y) // 10 20
	// fmt.Println(z) // undefined: z

	// but if we want to access the global variable a in the function printSomething
	printSomething() // 200 100
	// but if we want to access the local variable b in the main function
	// fmt.Println(b) // undefined: b
	// on the other hand if we want to access the global variable a in the main function
	fmt.Println(a) // 100


	// let us call the function addSomething

	s := 20
	t := 30
	result := addSomething(s, t)
	fmt.Println(result) // 50

	// fmt.Println(a, b) // undefined: b here b is not accessible in the main function, because b is a local variable of the function printSomething
 }
