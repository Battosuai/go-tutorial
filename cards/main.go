package main

func main() {
	// var card string = "Ace of Spades"
	// cards := newDeck()
	// cards = append(cards, "Six of Spades")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// cards.print()
	// fmt.Println(cards)

	cards := newDeck()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}
