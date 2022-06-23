package main

import "fmt"

func main() {
	explicit_conversion()
	arrayDeclaration()
}

func explicit_conversion() {
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) - y
	var d int = x + int(y)
	fmt.Println("Z =", z, "\nD =", d)
}

func arrayDeclaration() {
	var x [3]int // standard array initialized with zeros
	fmt.Println("Standard array with zeros: ", x)
	var y = [...]int{1, 2, 3}
	fmt.Println("Array created using ...", y)
	y[2] = 4
	fmt.Println("Changing variable with indexes", y)
	fmt.Println("Length func", len(y))
}
