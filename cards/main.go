package main

func main() {

	//cards := newDeckFromFile("yaron.cards")
	cards := NewDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("yaron.cards")
	cards.print()
	cards.shuffle()
	cards.print()
	// hand, remainingDeck := deal(cards, 3)
	// hand.print()
	// remainingDeck.print()

}
