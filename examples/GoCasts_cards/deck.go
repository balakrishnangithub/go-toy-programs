package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

// decksEqual compares two decks, returns true if equal else false.
// reflect.DeepEqual(d1, d2) is equal to decksEqual(d1, d2)
func decksEqual(d1, d2 deck) bool {
	if (d1 == nil) != (d2 == nil) {
		return false
	}
	if len(d1) != len(d2) {
		return false
	}
	for i, v := range d1 {
		if v != d2[i] {
			return false
		}
	}
	return true
}

// newDeck returns a new deck of 52 cards.
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"♠", "♥", "♦", "♣"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", cardValue, cardSuit))
		}
	}
	return cards
}

// deal splits deck into two.
// It returns deck of handSize and remaining deck.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// print prints the list of cards in deck.
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index+1, card)
	}
}

// toString returns deck as comma separated string.
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// saveToFile writes deck as comma separated string to file.
// It returns error of WriteFile.
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

// newDeckFromFile splits comma separated string from file.
// It returns []string as deck.
func newDeckFromFile(filename string) (deck, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	s := strings.Split(string(bytes), ",")
	return s, err
}

// shuffle shuffles the elements in deck
func (d deck) shuffle() {
	rand.Seed(time.Now().Unix())
	deckLength := len(d)
	for index := range d {
		randIndex := rand.Intn(deckLength) // (0,deckLength] greater than or equal to 0 and less than deckLength
		d[index], d[randIndex] = d[randIndex], d[index]
	}
}
