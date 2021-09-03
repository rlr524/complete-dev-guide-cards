package main

import (
	"fmt"
	"strings"
)

// Deck is a slice of strings
type Deck []string

func NewDeck() Deck {
	cards := Deck{}
	
	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardRankings := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}
	
	for _, suit := range cardSuits {
		for _, rank := range cardRankings {
			cards = append(cards, rank + " of " + suit)
		}
	}
	return cards
}

// Print is a receiver function of the Deck type (i.e. it extends Deck, thought we
// don't generally use that OOP-type terminology in Go...any variable of type Deck
// now has access to this print function), for example we use it in the main()
// function and access it with the cards variable, which is of type Deck
func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	deckSizeStart := d[:handSize]
	deckSizeEnd := d[handSize:]
	return deckSizeStart, deckSizeEnd
}

func (d Deck) toString() string {
	sliceAsString := strings.Join([]string(d), ",")
	return sliceAsString
}
