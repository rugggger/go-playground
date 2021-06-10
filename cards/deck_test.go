package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 40 {
		t.Errorf("Not enough cards")
	}

}
