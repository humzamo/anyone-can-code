package main

import "fmt"

func main() {

	// TASK 1.1: make a variable with your name and print out "Hello {name}"
	// Change the name and see what happens

	var name string
	name = "Humza"
	fmt.Println("Hello", name)

	// TASK 1.2: make a variable with your age and print out "My age is: {age}"

	age := 25
	fmt.Println("My age is:", age)

	// TASK 1.3: It's your birthday - add one to your age and print it out

	age = age + 1 // or you can do: age++
	fmt.Println("My new age is:", age)

	// TASK 1.4: make a boolean variable called "canDrive" to say if you can drive

	var canDrive bool
	canDrive = true
	fmt.Println("Can I drive?:", canDrive)
}
