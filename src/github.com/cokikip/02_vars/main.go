package main

import "fmt"

func main() {
	//Create a variable using var
	// var with variable type
	var name string = "Collins"
	//var without type
	var age = 26
	// using const you cannot reassign the variable
	const isCool = true
	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
}
