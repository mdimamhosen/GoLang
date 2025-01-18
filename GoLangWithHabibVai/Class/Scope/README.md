<!-- Scope in Go Lang -->

# Scope in Go

Scope in Go determines the visibility and lifetime of variables, constants, functions, and types. There are several types of scope in Go:

1. **Package Scope**: Identifiers declared at the package level are visible throughout the entire package.
2. **File Scope**: Identifiers declared within a file are visible only within that file.
3. **Function Scope**: Identifiers declared within a function are visible only within that function.
4. **Block Scope**: Identifiers declared within a block (e.g., within an `if`, `for`, or `switch` statement) are visible only within that block.

## Examples

### Package Scope

```go
package main

import "fmt"

var packageVar = "I am a package-level variable"

func main() {
    fmt.Println(packageVar)
}
```

### File Scope

```go
// file1.go
package main

import "fmt"

var fileVar = "I am a file-level variable"

func main() {
    fmt.Println(fileVar)
}
```

```go
// file2.go
package main

import "fmt"

func anotherFunction() {
    // fmt.Println(fileVar) // This will cause an error because fileVar is not visible here
}
```

### Function Scope

```go
package main

import "fmt"

func main() {
    var functionVar = "I am a function-level variable"
    fmt.Println(functionVar)
}
```

### Block Scope

```go
package main

import "fmt"

func main() {
    if true {
        var blockVar = "I am a block-level variable"
        fmt.Println(blockVar)
    }
    // fmt.Println(blockVar) // This will cause an error because blockVar is not visible here
}
```
