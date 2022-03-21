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
	var canDrive bool
	canDrive = CanDrive(age)
	fmt.Println("Can I drive?:", canDrive)

	// Task 2.2: make a function which can say what car you have based on your age and if you can drive
	// Return the output as a string.
	//"No car" if you can't drive, "Volvo" if you're under 30, "Tesla" if you're 30+
	fmt.Println("Which car do I have?:", WhichCar(canDrive, age))

}

func CanDrive(age int) bool {
	if age > 17 {
		return true
	} else {
		return false
	}
}

func WhichCar(canDrive bool, age int) string {
	if canDrive == false {
		return "No car for you :("
	}

	if age < 30 {
		return "Volvo"
	}

	return "Tesla"
}
