package main

func main() {
	// cards := newDeck()

	// hand, remainingHand := deal(cards, 5)
	// hand.print()
	// remainingHand.print()

	cards := newDeck()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

}
