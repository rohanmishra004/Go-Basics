package main

import "fmt"

func main() {

	/**
	lets break down the below line -
	var - to create a new variable
	card - here card is the name of the variable
	string - this bounds that only string values can be assigned to this variable

	Since go is static type we cannot change the value of variables dynamically


	**/
	// var card string = "Ace of cards"
	card := "Ace of cards" //this also does the same thing as above but it tells the go compiler to figure the type of data stored in card and go does that by reading this :=. Go will infer types to a certain degree unlike java or C++

	//Also := is used only when we are defining new values and not during reassigning

	card = "Five of Diamonds"
	fmt.Println(card)
}
