package main

import (
	"fmt"
	"math"
)

/*
Type conversions
~ The expression T(v) converts the value v to the type T.
~ Some numeric conversions:
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
~ Or, put more simply:
i := 42
f := float64(i)
u := uint(f)
*/

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
