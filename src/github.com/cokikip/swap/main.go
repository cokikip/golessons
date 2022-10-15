package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func addSub(x, y int) (int, int) {
	a := x + y
	b := x - y
	return a, b
}

func main() {
	a, b := swap("Collins", "Kiplagat")
	fmt.Println(a, b)
	fmt.Println(addSub(9, 3))
}
