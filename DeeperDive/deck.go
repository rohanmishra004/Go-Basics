package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

//shuffle - to shuffle we will use math/rand function -  func Intn - Intn returns as an int , a non - negative pseudo-random number in [0,n)  from the default Source. It panics if n<=0

// since this is a receiver function we donot need to set any arguments or set any return type since this is only going to shuffle the deck
func (d deck) shuffle() {
	for i := range d {
		//generate random number between 0 and len of deck -1
		//here the source for rand will remain the same so everytime the rand will return the same type of output , so in order to resolve this we create our source and so now everytime the new random series is generated it will be different
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		newPosition := r.Intn(len(d) - 1)
		//swap i with new random position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
