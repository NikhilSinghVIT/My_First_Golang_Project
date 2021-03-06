package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected a deck of 16 cards, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected value of Ace of spades but got %v", d[0])
	}

	if (d[len(d)-1]) != "Three of Hearts" {
		t.Errorf("expected value of Three of Hearts but got %v", d[len(d)-1])
	}
}
