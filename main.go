package main

import "fmt"

func main() {
	//cardGame()
	var cards = []string{"complete", newCard()}
	cards = append(cards, "combination")

	//fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func cardGame() {
	//var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Two of Spades"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "diamonds"
}

func intCard() int {
	return 45
}
