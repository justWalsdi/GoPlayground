package main

import (
	"errors"
	"fmt"
)

// somewhat a class
type student struct {
	name   string
	sex    string
	height float32
}

// now age is int, can be returned inside of a function
type age int

func main() {
	explicitConversion()
	arrayDeclaration()
	slices()
	fibonacci_wrapper(15)
	errWrapper()
	maps()
	structExample()
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

func fibonacci_of(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci_of(n-1) + fibonacci_of(n-2)
}

func fibonacci_wrapper(amount int) {
	fmt.Println("Fibonacci sequence:")
	for i := 0; i < amount; i++ {
		fmt.Print(fibonacci_of(i), " ")
	}
	fmt.Print("\n")
}

func err_examplePlusPlus(numberBelowZero int) (int, error) {
	if numberBelowZero < 0 {
		return int(0), errors.New("Number is below zero!")
	}
	numberBelowZero++
	return numberBelowZero, nil
}

func errWrapper() {
	var result int
	var err error

	result, err = err_examplePlusPlus(10)
	if err != nil {
		panic("Well the number is in fact below zero")
	}
	fmt.Printf("Result 1 is: %d\n", result)

	result, err = err_examplePlusPlus(-10)
	if err != nil {
		//panic("Well the number is in fact below zero")
	}
	//fmt.Printf("Result 2 is: %d\n", result)

	fmt.Println("This statement can't be reached.")
}

func maps() {
	//var mapName map[keyType]valueType
	var justSomeMap map[string]string
	justSomeMap = make(map[string]string, 2)
	justSomeMap["Hello"] = "world"
	justSomeMap["Hey"] = "partner"
	fmt.Println("Maps:", justSomeMap)

	justWalrusMap := make(map[int]string, 2)
	justWalrusMap[324] = "well well well"
	justWalrusMap[283] = "what the fuck is this"
	fmt.Println("Maps 2:", justWalrusMap)

	justMap := map[string]string{
		"hey":   "tony",
		"hello": "dad",
		"heyyy": "drunk",
	}
	fmt.Println("Maps 3:", justMap)

	example, exists := justMap["heyf"]
	fmt.Println("Example heyf:", example, "\nIs it exist:", exists)

	delete(justMap, "heyyy")
	fmt.Println("Maps 4, deleted elem heyyy:", justMap)
}

func structExample() {
	human := struct {
		name   string
		age    int
		salary uint
		height float32
	}{
		name:   "garry",
		age:    18,
		salary: 18000,
		height: 1.71,
	}

	fmt.Println("Human type:", human)

	studentExample := student{
		name:   "Josh",
		sex:    "M",
		height: 1.71,
	}

	fmt.Println("Student Josh:", studentExample)
	studentExample.printInfo()
	studentExample.setName("Mark").printInfo()
}

func (s student) printInfo() student {
	fmt.Printf("Name: %s\nSex: %s\nHeight: %.2f\n", s.name, s.sex, s.height)
	return s
}

func (s *student) setName(name string) *student {
	if name != "" {
		s.name = name
	}
	return s
}
