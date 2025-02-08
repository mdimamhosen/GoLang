package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// a and b are parameters

func main() {
	var sum = add(1, 2)
	fmt.Println(sum)
	// 1 and 2 are arguments
}
