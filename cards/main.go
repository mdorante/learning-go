package main

import "fmt"

// variables can also be assigned outside of a func, it just has to be done with long syntax
var decks int = 52

func main() {
	// creates a variable, explicitly specifying the data type
	var card string = "Ace of Spades"
	fmt.Println(card)

	// also creates a variable, but the Go compiler infers the data type from whatever value you're assigning to the variable
	card2 := "Five of Diamonds"
	fmt.Println(card2)

	// NOTE: ":=" is only used at variable initialization, if you want to reassign another value, you can use "="
	card2 = "Queen of Hearts"
	fmt.Println(card2)

	// we can also declare a variable without assigning a value to it, it can be assigned later
	var deckSize int
	deckSize = 20
	fmt.Println(deckSize)

	fmt.Println(decks)
	fmt.Println(hi)
}
