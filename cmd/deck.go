package main

import "fmt"

// Deck is a slice of strings
type Deck []string

// print is a receiver function of the Deck type (i.e. it extends Deck, thought we
// don't generally use that OOP-type terminology in Go...any variable of type Deck
// now has access to this print function)
func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
