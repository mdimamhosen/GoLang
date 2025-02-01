# Arrays in Go

## Declaring Arrays

You can declare an array in Go using the following syntax:

```go
var arrayName [size]elementType
```

Example:

```go
var numbers [5]int
```

## Initializing Arrays

Arrays can be initialized at the time of declaration:

```go
var numbers = [5]int{1, 2, 3, 4, 5}
```

Or you can use the shorthand notation:

```go
numbers := [5]int{1, 2, 3, 4, 5}
```

## Accessing Array Elements

Array elements are accessed using the index, which starts from 0:

```go
fmt.Println(numbers[0]) // Output: 1
```

## Iterating Over Arrays

You can iterate over arrays using a `for` loop:

```go
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

Or using the `range` keyword:

```go
for index, value := range numbers {
    fmt.Println(index, value)
}
```

## Multidimensional Arrays

Go supports multidimensional arrays. A two-dimensional array is declared as follows:

```go
var matrix [3][3]int
```

Example:

```go
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```

## Array of Arrays

You can also create an array of arrays:

```go
var arrayOfArrays [2][3]int
```

Example:

```go
arrayOfArrays := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

## Passing Arrays to Functions

Arrays can be passed to functions by value, meaning the function receives a copy of the array:

```go
func printArray(arr [5]int) {
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
}
```

To modify the original array, you can pass a pointer to the array:

```go
func modifyArray(arr *[5]int) {
    arr[0] = 100
}
```
