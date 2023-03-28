package main

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
	cards.saveToFile("my_cards")
	// cards := newDeckFromFile("not_existed_file_name")
	// fmt.Println(cards)
	// hand, remainingCards := deal(cards, 5)
	// remainingCards.print()
}
