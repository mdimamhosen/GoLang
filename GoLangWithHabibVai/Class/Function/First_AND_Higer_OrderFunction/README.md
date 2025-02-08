## First-Order Function and Higher-Order Function

### First-Order Function

A first-order function is a function that does not take other functions as arguments and does not return a function as its result. It operates on basic data types and structures.

### Higher-Order Function

A higher-order function is a function that can take other functions as arguments and/or return a function as its result. Higher-order functions are a key feature of functional programming languages like Haskell and Racket.

### Example Code

#### First-Order Function

```go
package main

import "fmt"

// First-order function: does not take or return another function
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println("Result:", result) // Output: Result: 8
}
```

#### Higher-Order Function

```go
package main

import "fmt"

// Higher-order function: takes a function as an argument
func applyOperation(a int, b int, operation func(int, int) int) int {
    return operation(a, b)
}

// Function to be passed as an argument
func multiply(a int, b int) int {
    return a * b
}

func main() {
    result := applyOperation(5, 3, multiply)
    fmt.Println("Result:", result) // Output: Result: 15
}
```

### Logic in Mathematics

In discrete mathematics, logic is used to define and analyze the properties and relationships of objects.

1. **Object**: An entity that has a physical existence (e.g., people, animals).
2. **Property**: Characteristics or attributes of objects (e.g., color, height).
3. **Relation**: Describes how objects are related to each other (e.g., "all customers must pay their pizza bills").

Example:

- **Object**: Customer
- **Property**: Has a bill
- **Relation**: Must pay the bill

- Rules: These are called First Order Logic. First order logic works with object, property and relation.

- If, we say:- Any rules that applies to all customer also applied to the tutul. Here this rules works with first order rules. So, it's called Higer Order Logic.

In the context of functions:

- **First-Order Function**: Operates directly on objects and their properties.
- **Higher-Order Function**: Operates on relations between functions, allowing more abstract and flexible operations.

### Functional Paradigms

Functional programming is a programming paradigm where programs are constructed by applying and composing functions. It is derived from mathematical functions and emphasizes the use of pure functions, immutability, and higher-order functions.

Key concepts in functional programming include:

- **Pure Functions**: Functions that always produce the same output for the same input and have no side effects.
- **Immutability**: Data cannot be changed once it is created. Instead, new data structures are created with the updated values.
- **First-Class Functions**: Functions are treated as first-class citizens, meaning they can be assigned to variables, passed as arguments, and returned from other functions.
- **Higher-Order Functions**: Functions that take other functions as arguments or return them as results.

Functional programming languages include Haskell, Racket, and Lisp, among others. These languages provide powerful abstractions for working with functions and enable developers to write concise and expressive code.

### Additional Example Code

#### Higher-Order Function Returning Another Function

```go
package main

import "fmt"

/*

Rules for Higher Order Functions:
1. Function can be passed as an argument to another function.
2. Function can return another function.
3. Function can be assigned to a variable.
4. Function can be stored in a data structure.

*/

func call() func(int, int) {
    return add
}

func add(a, b int) {
    z := a + b
    fmt.Println(z)
}

func main() {
    // Here call() is a higher-order function which is returning another function add.
    // The returned function is assigned to a variable f, then f is called with arguments 10, 20.
    f := call()
    f(10, 20) // Output: 30
}
```

#### Higher-Order Function with First-Class Functions

```go
package main

import "fmt"

// -> First Order Function is not First class function.
// -> Higher Order Function is First class function.

/*

Rules for Higher Order Functions:
1. Function can be passed as an argument to another function.
2. Function can return another function.
3. Function can be assigned to a variable.
4. Function can be stored in a data structure.

*/

func call() func(int, int) {
    return add
}

func add(a, b int) {
    z := a + b
    fmt.Println(z)
}

func main() {
    // Here call() is a higher-order function which is returning another function add.
    // The returned function is assigned to a variable f, then f is called with arguments 10, 20.
    f := call()
    f(10, 20) // Output: 30
}

// => First class citizen => that variables or functions can be passed as arguments to functions, returned from functions, and assigned to variables.
```
