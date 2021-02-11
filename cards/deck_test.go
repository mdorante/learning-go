package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	dl := len(d)
	if dl != 52 {
		// this will print out an error message if the test fails
		// %v is a placeholder, it will be substituted with the value of dl
		t.Errorf("Expected deck length of 52, but got %v", dl)
	}
}
