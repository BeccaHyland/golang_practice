package main

import "fmt"
// create a new type, "deck"
// which is a slice of strings
type deck []string

// after newDeck we state the type to be returned
// as we do with any function
// and in this case we are returning our own type we created
// which is "deck"
func newDeck deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}