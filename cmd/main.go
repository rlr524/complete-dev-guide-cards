package main

import "fmt"

const player string = "Madison"

func main() {
	cards := Deck{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

