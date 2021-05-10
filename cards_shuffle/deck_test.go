package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	var d []string = newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	fmt.Println("Test 1 passed")
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	fmt.Println("Test 2 passed")
	if d[len(d)-1] != "Jack of Clubs" {
		t.Errorf("Expected first card of Jack of Clubs, but got %v", d[len(d)-1])
	}
	fmt.Println("Test 3 passed")

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
	fmt.Println("Test 4 passed")
}
