package main

import "fmt"

func printValue(value int) {
	println(value)
}

func printValues(values ...int) {
	for _, value := range values {
		printValue(value)
	}

	fmt.Println(values)
}

func printDiffMultipleValues(strValue string, intValues ...int) {
	fmt.Println(strValue)
	fmt.Println(intValues)
}

func printMultiPleStringValues(strValues ...string) {
	fmt.Println(strValues)
}
func add(a, b int) int{
	return a+b
}
func main() {
	printValue(10)

	printValues(10, 20, 30)
	printDiffMultipleValues("Hello", 10, 20, 30)
	printMultiPleStringValues("Hello", "World", "Go")

	fmt.Println(add(10, 20))
}
