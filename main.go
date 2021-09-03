package main

import "fmt"

const player string = "Madison"

func main() {
	cards := NewDeck()
	cardsString := cards.toString()
	fmt.Println(cardsString)
}
