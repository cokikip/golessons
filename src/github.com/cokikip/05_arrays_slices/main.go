package main

import "fmt"

func main() {

	//Arrays
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "apple"
	fruitArr[1] = "mango"

	// Declare and assign arrays
	fruit := [2]string{"Mango", "Oranges"}
	fmt.Println(fruit)

	fmt.Println(fruitArr)

	fruitslice := []string{"Apple", "Orange", "Lemon", "Grape", "Citrus"}
	fmt.Println(fruitslice)
	fmt.Println(len(fruitslice))
	fmt.Printf("%T\n", fruitslice)

	// Get a range of data
	fmt.Println(fruitslice[1:3])
}
