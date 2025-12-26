package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
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

//Deck Deal

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Save to file - save a list of cards to a file on a local machine

func (d deck) toString() string {

	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

//load data from local file- just opposite of what we did above
//using readFile function

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		//option 1 - log the error and return a call to newDeck()

		//option 2 - log the error and exit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//now we have a bytesize and we want to convert it to slice of string , the exact opposite of what we were trying before with save on local storage
	s := strings.Split(string(bs), ",")
	return deck(s)
}
