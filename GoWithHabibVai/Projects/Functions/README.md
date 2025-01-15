# GoLang Functions Project

This project demonstrates various functions in GoLang, including addition, subtraction, multiplication, division, and printing messages.

## Functions

### add1

```go
func add1(x int, y int) {
	sum := x + y
	fmt.Println(sum)
}
```

- **Description**: Takes two integers, calculates their sum, and prints it.
- **Example**:

```go
add1(42, 13) // Output: 55
```

### add2

```go
func add2(x int, y int) int {
	return x + y
}
```

- **Description**: Takes two integers and returns their sum.
- **Example**:

```go
fmt.Println(add2(42, 13)) // Output: 55
```

### add_mul

```go
func add_mul(x int, y int) (int, int) {
	sum := x + y
	mul := x * y
	return sum, mul
}
```

- **Description**: Takes two integers and returns their sum and product.
- **Example**:

```go
sum, mul := add_mul(42, 13)
fmt.Println("Sum:", sum) // Output: Sum: 55
fmt.Println("Mul:", mul) // Output: Mul: 546
```

### printSomething

```go
func printSomething() {
	fmt.Println("Something is printed")
}
```

- **Description**: Prints a predefined message.
- **Example**:

```go
printSomething() // Output: Something is printed
```

### sayHello

```go
func sayHello(str string) {
	fmt.Println("Hello,", str)
}
```

- **Description**: Takes a string and prints a greeting message.
- **Example**:

```go
sayHello("World") // Output: Hello, World
```

### subtract

```go
func subtract(x int, y int) int {
	return x - y
}
```

- **Description**: Takes two integers and returns their difference.
- **Example**:

```go
fmt.Println("Difference:", subtract(42, 13)) // Output: Difference: 29
```

### divide

```go
func divide(x int, y int) float64 {
	if y != 0 {
		return float64(x) / float64(y)
	}
	fmt.Println("Error: Division by zero")
	return 0
}
```

- **Description**: Takes two integers and returns their quotient. Prints an error message if division by zero is attempted.
- **Example**:

```go
fmt.Println("Quotient:", divide(42, 13)) // Output: Quotient: 3.230769230769231
```

## Main Function

The `main` function demonstrates the usage of all the above functions.

```go
func main() {
	add1(42, 13)
	fmt.Println(add2(42, 13))
	sum, mul := add_mul(42, 13)
	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", mul)
	printSomething()
	sayHello("World")
	fmt.Println("Difference:", subtract(42, 13))
	fmt.Println("Quotient:", divide(42, 13))
}
```
