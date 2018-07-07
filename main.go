package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Ace Of Spades"
}
