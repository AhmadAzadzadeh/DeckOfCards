package main

func main() {
	cards := deck{"Ace Of Diamonds", newCard()}
	cards = append(cards, "Six Of Spades")
	cards.print()
}

func newCard() string {
	return "Ace Of Spades"
}
