package main

import "fmt"

func main() {
	a := 5
	b := &a // value of a memory address

	fmt.Println(a, b)
	fmt.Printf("%T %T\n", a, b)

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
