package main

import "testing"

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d)!=52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades"{
		t.Errorf("Expecting Ace of Spades as the first card but got %v",d[0])
	}

	if d[len(d)-1] != "King of Clubs"{
		t.Errorf("Expecting King of Clubs as the first card but got %v",d[0])
	}
}

