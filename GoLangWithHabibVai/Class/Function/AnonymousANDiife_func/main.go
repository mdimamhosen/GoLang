package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}

	func() {
		fmt.Println("Hello from immediately invoked function")
	} ()


	func (a, b int) {
		fmt.Println(a + b)
	}	(10, 20)

	fmt.Println(add(10, 20))
}
