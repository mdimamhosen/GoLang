package main

import "fmt"
func main() {
	fmt.Println("Hello, World! This is my first go program." )

	a := 10
	b := 20

	if(a > b) {
		fmt.Println("a is greater than b")
	} else if(a <= b) {
		fmt.Println("a is less than or equal to b")
	} else {
		fmt.Println("a is equal to b")
	}
}
