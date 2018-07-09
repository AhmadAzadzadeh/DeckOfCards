package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected Length of Deck is 52 but the length is: %v", len(d))
	}

	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected first card of Ace Of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King Of Clubs" {
		t.Errorf("Expected last card of King Of Clubs, but got %v", d[len(d)-1])
	}
}
