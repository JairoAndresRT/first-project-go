package main

import (
	"fmt"
)

func variables() {
	// First way
	var firstName, lastName string
	var age int

	// Second way
	var (
		firstName2, lastName2 string
		age2                  int
	)

	// Third way
	var (
		firstName3 = "Jairo3"
		lastName3  = "Romero3"
		age3       = 26
	)

	// Fourth way
	var firstName4, lastName4, age4 = "Jairo4", "Romero4", 26

	// Fifth way: We only can use this way where we are inside on a function
	firstName5, lastName5, age5 := "Jairo4", "Romero4", 26

	firstName = "Jairo"
	lastName = "Romero"
	age = 26
	firstName2 = "Jairo2"
	lastName2 = "Romero2"
	age2 = 26
	fmt.Println("1) First name: ", firstName, " Last name: ", lastName, " Age: ", age)
	fmt.Println("2) First name: ", firstName2, " Last name: ", lastName2, " Age: ", age2)
	fmt.Println("3) First name: ", firstName3, " Last name: ", lastName3, " Age: ", age3)
	fmt.Println("4) First name: ", firstName4, " Last name: ", lastName4, " Age: ", age4)
	fmt.Println("5) First name: ", firstName5, " Last name: ", lastName5, " Age: ", age5)
}
