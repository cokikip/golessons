package main

import "fmt"

var name string = "Collins"

func main() {
	//Create a variable using var
	// var with variable type
	// var name string = "Collins"

	// shorthand var
	size := 1.9993
	// name := "Collins"
	name, email := "Collins", "collins@gmail.com"

	//var without type
	var age = 26
	// using const you cannot reassign the variable
	const isCool = true
	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}
