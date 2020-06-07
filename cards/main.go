package main

func main() {
	newDeck().saveToFile("deckofcards")
	cards := newDeckFromFile("deckofcards")
	cards.shuffle()
	cards.print()
}
