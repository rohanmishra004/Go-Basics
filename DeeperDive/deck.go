package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	//in order to create the all 52 cobinations we are going to seperate the suits and values in 2 different sets and then use for loop twice to map the values to deck

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//after this we will setup two for loops nested inside one another to iterate through all the combinations

	//while looping if we have some variable which we are not going to use , we replace that variable with an _
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
