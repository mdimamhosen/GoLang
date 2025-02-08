## Difference between Parameter and Argument

### Parameter

A parameter is a variable in the declaration of a function. Parameters are used to define the inputs that a function can accept.

### Argument

An argument is the actual value that is passed to the function when it is called. Arguments are the real values that are supplied to the function's parameters.

### Example Code

```go
package main

import "fmt"

// Function with parameters 'a' and 'b'
func add(a int, b int) int {
    return a + b
}

func main() {
    // Calling the function with arguments 5 and 3
    result := add(5, 3)
    fmt.Println("Result:", result) // Output: Result: 8
}
```

In the above example:

- `a` and `b` are parameters of the `add` function.
- `5` and `3` are arguments passed to the `add` function when it is called in the `main` function.
