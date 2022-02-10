package main

import (
	"fmt"
)

// Create new type 'deck'

type deck []string 

func newDeck() deck {
	cards := deck{}
	
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardRank := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _,suit := range cardSuits {
		for _,rank := range cardRank {
			cards = append(cards, rank+" of "+suit)
		}
	}

	return cards
}

func (dk deck) print() {
	for i, card := range dk {
		fmt.Println(i, card)
	}
}

func deal(dk deck, handSize int) (deck, deck) {
	return dk[:handSize], dk[handSize:]
}

// Take a deck and return a string
func (dk deck) toString() string {
	str := ""
	for _, card := range dk {
		str = str + card + ","
	}
	return str
}

//Alternative
//func (dk deck) toString() string {
//	return strings.Join([]string(d), ",")
//}

