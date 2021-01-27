package main

func main() {
	cards := deck{"Queen of Hearts", newCard()}
	cards = append(cards, "Ace of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
