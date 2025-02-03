package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Deck is a slice of strings
type Deck []string

func newDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardRankings := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight",
		"Seven", "Six", "Five", "Four", "Three", "Two"}

	for _, suit := range cardSuits {
		for _, rank := range cardRankings {
			cards = append(cards, rank+" of "+suit)
		}
	}
	return cards
}

// print is a receiver function of the Deck type (i.e. it extends Deck, thought we
// don't generally use that OOP-type terminology in Go...any variable of type Deck
// now has access to this print function), for example we use it in the main()
// function and access it with the cards variable, which is of type Deck
func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	deckSizeStart := d[:handSize]
	deckSizeEnd := d[handSize:]
	return deckSizeStart, deckSizeEnd
}

func (d Deck) toString() string {
	sliceAsString := strings.Join(d, ",")
	return sliceAsString
}

func (d Deck) saveToFile(filename string) error {
	// Look at the file when it is created. Macs have a default mask of 022 on files that are
	// newly created. Basically, this means take the permissions that were passed in, like 777, and
	return os.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) Deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(string(bs), ",")
	return s
}
