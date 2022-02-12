package main

import "testing"

//t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length 52, but got %d", len(d))
	}
}