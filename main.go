package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	for i, card := range cards {
		fmt.Print(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
