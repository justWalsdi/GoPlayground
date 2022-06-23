package main

import (
	"fmt"
)

func main() {
	explicitConversion()
	arrayDeclaration()
	slices()
}

func explicitConversion() {
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

func slices() {
	var x = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice example:", x)
	var y []int
	fmt.Println("Y is an empty slice and equals to nil:", y == nil) // slices can't be compared to anything but nil
	// you can use reflect package and its function DeepEqual but it is intended for testing.
	y = append(y, 10, 10, 10) // to add elements use append and reassign
	y = append(y, x...)       // to add another slice use "slice_name..."
	fmt.Println("Appended elements to empty slice")
	fmt.Println("Capacity of y slice:", cap(y))

	// to create slices of specific size ([3]int would not work because it will create an array) use make
	// walrus operator(`:=`) is used to assign a type at the compile time
	z := make([]int, 5)
	fmt.Println("Empty zero initialized slice:", z, "size and cap:", len(z), cap(z))
	w := make([]int, 5, 10) // to make capacity 10 at the start
	fmt.Println("Empty zero initialized slice:", w, "size and cap are different:", len(w), cap(w))
}
