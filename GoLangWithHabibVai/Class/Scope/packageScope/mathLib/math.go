package mathLib

import "fmt"

// Sum should be capitalized to be exported
// Like if you want to use this function in another package you should capitalize it.
func Sum(n1, n2 int)   {
	fmt.Print(n1 + n2)
}
