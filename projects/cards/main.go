package main

import "log"

func main() {
	newDeck().saveToFile("deckofcards")
	cards, err := newDeckFromFile("deckofcards")
	if err != nil {
		log.Fatal(err)
	}
	cards.shuffle()
	cards.print()
}
