package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// test deck length
	dl := len(d)
	if dl != 52 {
		// this will print out an error message if the test fails
		// %v is a placeholder, it will be substituted with the value of dl
		t.Errorf("Expected deck length of 52, but got %v", dl)
	}

	// test first card
	fc := d[0]
	if fc != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", fc)
	}

	// test last card
	lc := d[len(d)-1]
	if lc != "King of Diamonds" {
		t.Errorf("Expected last card to be King of Diamonds, but got %v", lc)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// delete file named _deckTestFile
	os.Remove("_deckTestFile")

	d := newDeck()

	d.saveToFile("_deckTestFile")

	loadedDeck := newDeckFromFile("_deckTestFile")

	dl := len(loadedDeck)

	if dl != 52 {
		t.Errorf("Expected deck length of 52, but got %v", dl)
	}

	os.Remove("_deckTestFile")
}
