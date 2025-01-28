package main

import "fmt"

func main() {

	age := 20

	if age < 30 {
		age := 25
		fmt.Print(age)
	}

	fmt.Print(age)
}
