package main

func main() {
	// Function
	func() {
		println("Function")
	}()

	// Function Expression or Assigning a function to a variable
	f := func() {
		println("Function Expression")
	}
	f()

	// Local function expression

	globalAdd(10, 20) // works fine because globalAdd is defined before this line and it is a global function

	// summation(10, 20) // This will not work because summation is defined after this line and it is a local function

	summation := func(a, b int) {
		println(a + b)
	}

	summation(10, 20)

}

func init() {
	println("Init function")
}

func globalAdd(a, b int) {
	println(a + b)
}
