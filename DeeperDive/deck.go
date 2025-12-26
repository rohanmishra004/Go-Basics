package main

import "fmt"

/**
Create a new type of deck which is a slice of strings which extends the functionality of slice
**/

type deck []string

//here we want to write a new function to print all the values of the cards

//here (d deck) is referred to as the receiver

/**
func (d deck) print(){
}

Any variable of type deck now gets access to the print method.

in the main file we will see that while assigning we have assigned cards to be of type deck. therefore cards will have access to all the functions that are assigned to
**/

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
