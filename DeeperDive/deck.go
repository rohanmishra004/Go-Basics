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

//in this function the deal function is returning two values - both of type deck

// deal function takes two arguments - deck and handSize which is of int type
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Save to file - save a list of cards to a file on a local machine

/**
we are going to use io util package to help with  interacting with  the local machine

We will be using io util writeFile function -
Syntax -
func WriteFile(filename string, data []byte, perm fs.FileMode) error


since writeFile require data to be in byte slice format - we can use typeconversions to convert deck type to string and then byte

SYNTAX - []byte("hi there!")

currently we have deck(which is of type deck) --> []string --> string --> []byte(we want)
We need to follow the above conversion method

its better to write this save to file function inside of a helper file

**/

// this is a receiver function to convert the cards deck type to string
func (d deck) toString() string {
	//since our deck is basically a slice of string , we can condence this slice of string and convert it down to string
	return strings.Join([]string(d), ",")
	//we will use join function which is part of strings package and condence this slice of strings - func Join( a[]string, sep string) string --> return type is string

}

//Now since we have already converted the deck type to string all that is remaining is to save this data on to local machine

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
