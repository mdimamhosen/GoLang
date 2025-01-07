# Go Programming Notes 1

## Packages

- Every Go program must have a **`package`** declaration at the top.
- The entry point of every Go program is the **`main`** package.

### Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

To run the code, type

````
go run filename
               ```
## Importing Packages

- The **`import`** statement is used to include external or standard libraries in your program.
- For example:
  ```go
  import "fmt"
````

- `fmt` is short for **format**, used for formatted I/O operations like printing to the console.

---

## Functions

- A **function** in Go is declared using the **`func`** keyword.
- The main function has the signature:
  ```go
  func main() {
      // Code here
  }
  ```

### Example:

```go
func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(5, 3)
    fmt.Println("Sum:", sum)
}
```

---

## Variables

### Declaration and Initialization

1. **Short Declaration**:

   ```go
   a := 10
   b := "Hello"
   ```

   - The type is inferred automatically.

2. **Explicit Declaration**:

   ```go
   var a int = 10
   var b string = "Hello"
   ```

3. **Multiple Declaration**:

   ```go
   var (
       a int = 10
       b string = "Hello"
       c bool = true
   )
   ```

4. **Constant Declaration**:
   ```go
   const pi = 3.14
   ```

---

## Condition Statements

### `if-else`

- Syntax:

  ```go
  if condition {
      // Code to execute if condition is true
  } else if anotherCondition {
      // Code to execute if anotherCondition is true
  } else {
      // Code to execute if none of the conditions are true
  }
  ```

- Example:

  ```go
  func main() {
      a := 10

      if a > 5 {
          fmt.Println("a is greater than 5")
      } else if a == 5 {
          fmt.Println("a is equal to 5")
      } else {
          fmt.Println("a is less than 5")
      }
  }
  ```

### `switch`

- Syntax:

  ```go
  switch variable {
  case value1:
      // Code to execute
  case value2:
      // Code to execute
  default:
      // Code to execute if no cases match
  }
  ```

- Example:

  ```go
  func main() {
      day := 3

      switch day {
      case 1:
          fmt.Println("Monday")
      case 2:
          fmt.Println("Tuesday")
      case 3:
          fmt.Println("Wednesday")
      default:
          fmt.Println("Invalid day")
      }
  }
  ```

---

## Loops

### `for` Loop

- Go does not have a `while` loop; you can use `for` instead.
- Syntax:

  ```go
  for initialization; condition; increment {

  }
  ```

- Example:
  ```go
  func main() {
      for i := 1; i <= 5; i++ {
          fmt.Println("i:", i)
      }
  }
  ```

### Infinite Loop

- Use `for` without any conditions:
  ```go
  for {
      fmt.Println("This is an infinite loop")
  }
  ```

### `while` Equivalent

- Using `for` as a `while` loop:
  ```go
  func main() {
      i := 1
      for i <= 5 {
          fmt.Println("i:", i)
          i++
      }
  }
  ```

### Loop with `break` and `continue`

- Example:
  ```go
  func main() {
      for i := 1; i <= 10; i++ {
          if i == 5 {
              continue
          }
          if i == 8 {
              break
          }
          fmt.Println("i:", i)
      }
  }
  ```
