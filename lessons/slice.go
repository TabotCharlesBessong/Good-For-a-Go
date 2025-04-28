package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, "Good morning")
	// fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i,card)
	}
}

func newCard() string {
	return "Hello"
}