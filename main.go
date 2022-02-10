package main

import "fmt"

func main(){
	// var card string = "Ace of Spades"
	card := "Ace of Spades"

	fmt.Println(card)
	
	cards := deck{newCard(), newCard(), card, "Ace of Diamonds"}
	
	fmt.Println(cards)

	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	for i := 0; i < len(cards); i++ {
		fmt.Println(i, cards[i])
	}

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}