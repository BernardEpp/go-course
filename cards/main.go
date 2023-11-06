package main

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// fmt.Println(cards)
	cards := newDeckFromFile("myCards")
	cards.shuffle()
	cards.print()

}
