package main

import "fmt"

func main() {
	fmt.Print("Hello World from main function")
}

func init() {
	fmt.Println("Hello World from init function")
}
// init function is called before the main function. It is used to initialize the variables or to perform any task before the main function is called. It is only called only after the file execution is completed. First, the init function is called and then the main function is called.
