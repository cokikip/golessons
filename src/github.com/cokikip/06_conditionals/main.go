package main

import "fmt"

func main() {
	x := 55
	y := 10
	color := "blue"
	// %b is for base 2: binary
	// %d is for base 10:decimal

	if x < y {
		fmt.Printf("%b is less than %d\n", x, y)

	} else {
		fmt.Printf("%d is greator than %d\n", x, y)
	}

	if color == "red" {
		println("The color is red")
	} else if color == "blue" {
		println("The color is blue")
	} else {
		println("Other colors")
	}

	fmt.Println("Hello world")
}
