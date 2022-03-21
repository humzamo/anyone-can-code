package main

import "fmt"

func main() {
	var name string
	name = "Humza"
	fmt.Println("Hello", name)

	age := 25
	fmt.Println("My age is:", age)

	// TASK 2.1: make a function which says if you can drive based on your age
	// Assume that if you're older than 17, you can drive

	// Task 2.2: make a function which can say what car you have based on your age and if you can drive
	// Return the output as a string.
	// "No car" if you can't drive, "Volvo" if you're under 30, "Tesla" if you're 30
}
