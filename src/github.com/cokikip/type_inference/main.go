package main

import "fmt"

/*
~When declaring a variable without specifying an explicit type
(either by using the := syntax or var = expression syntax),
the variable's type is inferred from the value on the right hand side.
*/

func main() {
	v := 42    // int
	b := 42.23 // float64
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("b is of type %T\n", b)
}
