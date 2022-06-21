package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	explicit_conversion()
}

func explicit_conversion() {
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) - y
	var d int = x + int(y)
	fmt.Println("Z =", z, "\nD =", d)
}
