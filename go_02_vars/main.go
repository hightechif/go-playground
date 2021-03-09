package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Dhan" // or var name string = "Dhan"
	var age = 37      // or var age int = 37
	var isCool = true // You can use const like ES6
	isCool = false
	var size float32 = 2.3

	// Shorthand
	// name := "Dhan"
	// email := "person@email.com"
	// size := 1.3

	// Another way
	name, email := "Dhan", "person@email.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", email)
}
