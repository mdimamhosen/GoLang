# Function Expressions and Anonymous Functions in Go

In Go, functions can be assigned to variables, passed as arguments to other functions, and returned from functions. These are known as function expressions. Additionally, Go supports anonymous functions, which are functions without a name.

## Function Expressions

A function expression is simply a function assigned to a variable.

```go
package main
import "fmt"

func main() {
    add := func(a int, b int) int {
        return a + b
    }

    result := add(3, 4)
    fmt.Println("Sum:", result)
}
```

## Anonymous Functions

Anonymous functions are functions that are defined without a name. They are often used for short-lived operations.

```go
package main
import "fmt"

func main() {
    func(message string) {
        fmt.Println(message)
    }("Hello, Go!")
}
```

## Immediately Invoked Function Expressions (IIFE)

In Go, you can also create and immediately invoke anonymous functions. This is useful for initializing variables or executing code that doesn't need to be reused.

```go
package main
import "fmt"

func main() {
    result := func(a int, b int) int {
        return a + b
    }(3, 4)

    fmt.Println("Sum:", result)
}
```
