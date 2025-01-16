# Operators are used to perform operations on variables and values

## Arithmetic Operators

### 1. Addition: The + operator adds two operands. For example, x+y.

### 2. Subtraction: The - operator subtracts two operands. For example, x-y.

### 3. Multiplication: The * operator multiplies two operands. For example, x*y.

### 4. Division: The / operator divides the first operand by the second. For example, x/y.

### 5. Modulus: Returns the division remainder

### 6. Increment: The ++ Increases the value of a variable by 1

### 7. Decrement: The -- Decreases the value of a variable by 1

```go

package main
import "fmt"

func main() {
	fmt.Println(2 + 2) // 4
	fmt.Println(2 - 2) // 0
	fmt.Println(2 * 2) // 4
	fmt.Println(2 / 2) // 1
	fmt.Println(2 % 2) // 0
}
```

### Increment

```go

package main
import ("fmt")

func main() {
  x:= 10
  x++ // Add one new value (increment)
  fmt.Println(x)
}
```

### Decrement

```go

package main
import ("fmt")

func main() {
  x:= 10
  x-- // Remove one new value (decrement)
  fmt.Println(x)
}
```

## Assignment Operators

### Assignment operators are used to assign values to variables.

| Operator | Example | Same As    |
| -------- | ------- | ---------- |
| =        | x = 5   | x = 5      |
| +=       | x += 3  | x = x + 3  |
| -=       | x -= 3  | x = x - 3  |
| \*=      | x \*= 3 | x = x \* 3 |
| /=       | x /= 3  | x = x / 3  |
| %=       | x %= 3  | x = x % 3  |

## Questions and Examples

### 1. Create a variable named `lgNumber`, store 1000 as a value, and perform various operations.

```go
package main
import ("fmt")

func main() {
  lgNumber := 1000

  // Add that variable with itself
  sum := lgNumber + lgNumber
  fmt.Println("Addition:", sum) // 2000

  // Subtract that variable by itself
  difference := lgNumber - lgNumber
  fmt.Println("Subtraction:", difference) // 0

  // Multiply that variable with itself
  product := lgNumber * lgNumber
  fmt.Println("Multiplication:", product) // 1000000

  // Divide that variable with itself
  quotient := lgNumber / lgNumber
  fmt.Println("Division:", quotient) // 1
}
```
