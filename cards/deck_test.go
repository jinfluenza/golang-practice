package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	d := newDeck()
	d.saveToFile("_decktesting.txt")

	loadedDeck := newDeckFromFile("_decktesting.txt")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %d", len(loadedDeck))
	}

	os.Remove("_decktesting.txt")
}
