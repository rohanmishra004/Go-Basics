package main

func main() {
	// cards := newDeck()

	// hand, remainingHand := deal(cards, 5)
	// hand.print()
	// remainingHand.print()

	// cards := newDeckFromFile("my")
	// cards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
