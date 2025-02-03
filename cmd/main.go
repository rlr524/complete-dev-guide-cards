package main

const player string = "Madison"

func main() {
	//cards := newDeck()
	//cardsString := cards.toString()
	//fmt.Println(cardsString)
	//
	//err := cards.saveToFile("madi_cards")
	//if err != nil {
	//	return
	//}

	cards := newDeckFromFile("madi_cards")
	cards.print()
}
