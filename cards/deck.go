package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	// converts the deck to a slice of strings, then joins it into a single string, separated by ","
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0600)
}

func newDeckFromFile(filename string) deck {
	// returns byte slice and potential error object if the read goes wrong
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	// turns the byte slice into a string and the splits it into a string slice
	s := strings.Split(string(bs), ",")

	//  returns the string slice converted to a deck (which is just an "extended" string slice)
	return deck(s)
}
