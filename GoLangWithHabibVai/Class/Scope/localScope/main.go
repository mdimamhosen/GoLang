package main

import "fmt"



func main() {

	 x := 10

	 fmt.Print(x) // 10

	 if x >= 10 {
		 y := 20
		 fmt.Print(y) // 20
		//  fmt.Print(z) // undefined: z
	 }
	//  fmt.Print(y) // undefined: y because y is a local variable of the if block
 }
