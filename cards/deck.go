package main

import "fmt"

// Create a new 'deck' type, which is like an "extension" of a slice of strings
type deck []string

// Create a "receiver function", which is like a method for all variables of type "deck"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
