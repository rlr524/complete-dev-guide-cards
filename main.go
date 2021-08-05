package main

const player string = "Madison"

func main() {
	cards := NewDeck()
	hand, remainingCards := Deal(cards, 5)
	hand.Print()
	remainingCards.Print()
}
