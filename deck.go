package main

import "fmt"

// Create new type 'deck'

type deck []string 

func (dk deck) print() {
	for i, card := range dk {
		fmt.Println(i, card)
	}
}
