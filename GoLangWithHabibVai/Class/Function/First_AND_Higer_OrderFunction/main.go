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

func call() func (int, int)   {
    return add
}

func add(a, b int) {
    z := a + b

    fmt.Println(z)
}
func main() {
    // Here call() is a h o f which is returning another function add which is passed as an argument to call() function and then call() function is called after assigning it to a variable f then f is called with arguments 10, 20.
    f := call()
    f(10, 20)
}


// => First class citizen =>  that varibales or functions can be passed as arguments to functions, returned from functions, and assigned to variables.



