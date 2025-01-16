# GoLang Strings

This document provides an overview of string operations in GoLang with examples.

## Basic String Operations

### Concatenation

You can concatenate strings using the `+` operator.

```go
fmt.Println("Hello, " + "World!")
```

### Formatting Strings

Go provides several ways to format strings using `fmt.Printf`.

```go
message := "Hello World!"
fmt.Printf("Data: %v\n", message)
fmt.Printf("Data: %+v\n", message)
fmt.Printf("Data: %#v\n", message)
fmt.Printf("Data: %T\n", message)
fmt.Printf("Data: %q\n", message)
fmt.Printf("First character: %c\n", message[0])
```

### String Length

You can get the length of a string using the `len` function.

```go
fmt.Println("Length:", len(message))
```

### Substrings

You can extract substrings using slice notation.

```go
fmt.Println("Substring:", message[0:5]) // this will print "Hello"
```

### String Comparison

You can compare strings using `==` and `!=` operators or the `strings.Compare` function.

```go
msg1 := "one"
msg2 := "two"

fmt.Println("Equal?", msg1 == msg2)
fmt.Println("Not Equal?", msg1 != msg2)
fmt.Println(strings.Compare(msg1, msg2))
```

## Additional String Functions

### Contains

Check if a string contains a substring.

```go
fmt.Println(strings.Contains(message, "World")) // true
```

### ToUpper and ToLower

Convert strings to upper or lower case.

```go
fmt.Println(strings.ToUpper(message)) // "HELLO WORLD!"
fmt.Println(strings.ToLower(message)) // "hello world!"
```

### Split

Split a string into a slice of substrings.

```go
fmt.Println(strings.Split(message, " ")) // ["Hello", "World!"]
```

### Replace

Replace occurrences of a substring.

```go
fmt.Println(strings.Replace(message, "World", "Go", 1)) // "Hello Go!"
```

### Trim

Trim leading and trailing spaces.

```go
fmt.Println(strings.TrimSpace("  Hello World!  ")) // "Hello World!"
```

Refer to the [Go documentation](https://golang.org/pkg/strings/) for more string functions and detailed usage.
