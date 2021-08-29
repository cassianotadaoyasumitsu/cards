package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards { // i = index
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
