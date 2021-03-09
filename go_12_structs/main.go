package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (ponter receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Medina", lastName: "Putri", city: "Bandung", gender: "f", age: 22}
	// Alternative
	person2 := Person{"Ridhan", "Fadhilah", "Bandung", "m", 24}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	person1.hasBirthday()
	person1.getMarried("Fadhilah")
	person2.getMarried("Putri")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
