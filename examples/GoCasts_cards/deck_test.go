package main

import (
	"os"
	"reflect"
	"testing"
)

// Test_decksEqual is bootstrapped by https://github.com/cweill/gotests
func Test_decksEqual(t *testing.T) {
	d := newDeck()
	deckClone := make(deck, len(d))
	copy(deckClone, d)

	deckCloneMutated := make(deck, len(d))
	copy(deckCloneMutated, d)
	deckCloneMutated[0] = ""

	type args struct {
		d1 deck
		d2 deck
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "d1=nil, d2=nil", args: args{d1: nil, d2: nil}, want: reflect.DeepEqual(nil, nil)},
		{name: "d1=newDeck(), d2=nil", args: args{d1: d, d2: nil}, want: reflect.DeepEqual(d, nil)},
		{name: "d1=newDeck(), d2=newDeck()[len(d)/2:]", args: args{d1: d, d2: d[len(d)/2:]}, want: reflect.DeepEqual(d, d[len(d)/2:])},
		{name: "d1=newDeck(), d2=newDeck()", args: args{d1: d, d2: deckClone}, want: reflect.DeepEqual(d, deckClone)},
		{name: "d1=newDeck(), d2=newDeck() mutated", args: args{d1: d, d2: deckCloneMutated}, want: reflect.DeepEqual(d, deckCloneMutated)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decksEqual(tt.args.d1, tt.args.d2); got != tt.want {
				t.Errorf("decksEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecksEqual(t *testing.T) {
	if !decksEqual(nil, nil) {
		t.Errorf("Expected true for comparing two nil decks.")
	}
	if decksEqual(newDeck(), nil) {
		t.Errorf("Expected false for comparing not nil deck with a nil deck.")
	}
	d := newDeck()
	deckClone := make(deck, len(d))
	copy(deckClone, d)
	if decksEqual(d, d[len(d)/2:]) {
		t.Errorf("Expected false for comparing decks of different length.")
	}
	if !decksEqual(d, deckClone) {
		t.Errorf("Expected true for comparing cloned deck.")
	}
	deckClone[0] = ""
	if decksEqual(d, deckClone) {
		t.Errorf("Expected false for comparing two decks of same length with different values.")
	}
}

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v.", len(d))
	}
	if d[0] != "Ace of ♠" {
		t.Errorf(`Expected first card of "Ace of ♠", but got "%v".`, d[0])
	}
	if d[len(d)-1] != "King of ♣" {
		t.Errorf(`Expected first card of "Ace of ♠", but got "%v".`, d[0])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	inHand, remaining := deal(d, len(d)/2)
	if len(inHand) != len(d)/2 {
		t.Errorf("Expected inHand deck of length %v, but got %v.", len(d)/2, len(inHand))
	}
	if len(remaining) != len(d)/2 {
		t.Errorf("Expected remaining deck of length %v, but got %v.", len(d)/2, len(remaining))
	}
	inHandPlusRemaining := append(inHand, remaining...)
	if !decksEqual(d, inHandPlusRemaining) {
		t.Errorf("Expected inHand plus remaining deck to be same as initial deck.")
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	defer os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)

	loadedDeck, err := newDeckFromFile(filename)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !decksEqual(d, loadedDeck) {
		t.Errorf("Expected decks to be equal for saving and loading with same file name.")
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	deckClone := make(deck, len(d))
	copy(deckClone, d)
	deckClone.shuffle()
	if decksEqual(d, deckClone) {
		t.Errorf("Expected decks to be different after shuffle.")
	}
}
