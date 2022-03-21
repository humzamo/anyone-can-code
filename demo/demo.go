package main

import "fmt"

func main() {
	/* VARIABLES */

	// words/letters are of type "string"
	// numbers are of type "int"
	// yes/no or true/false are of type "bool"

	// how to declare
	// first way
	// - use the syntax: var NameOfVariable TypeOfVariable
	// - to assign a value: NameOfVariable = ValueYouWant

	var favouriteFood string
	favouriteFood = "burgers"
	fmt.Println("my favourite food is", favouriteFood)

	// second way
	// use the syntax: NameOfVariable := ValueYouWant
	// Golang works out the type for you ;)

	numberOfPizzas := 3
	fmt.Println("How many pizzas did I eat?", numberOfPizzas)

	likesPizza := true
	fmt.Println("Do I like pizza?", likesPizza)

	/* FUNCTIONS */

	// This is a template of a function
	// func NameOfYourFunction(input INPUT-TYPE) OUTPUT-TYPE{
	// Do something here with you INPUT-TYPE
	// return something-with-the-output-type
	//}

	favouriteFood = WhatsMyFavouriteFood()
	fmt.Println("my favourite food is", favouriteFood)
	numberOfPizzas = EatPizza(numberOfPizzas)

	/* IF/ELSE */

	// if condition {
	// 	// Do something
	// } else {
	// 	// Do something else
	// }

	tooManyPizzas := DidIEatTooManyPizzas(numberOfPizzas)
	fmt.Println("Did I eat many pizzas?", tooManyPizzas)
}

func WhatsMyFavouriteFood() string {
	return "pasta"
}

func EatPizza(number int) int {
	return number + 1
}

func DidIEatTooManyPizzas(number int) bool {
	if number > 5 {
		return true
	} else {
		return false
	}
}
