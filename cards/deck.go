package main

import "fmt"

// Create a new 'deck' type, which is like an "extension" of a slice of strings
type deck []string

// Creates a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// using '_' for index vars, let's Go know that I don't need to use these vars
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Create a "receiver function", which is like a method for all variables of type "deck"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deals a hand of cards, also returns the remaining deck
func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}
