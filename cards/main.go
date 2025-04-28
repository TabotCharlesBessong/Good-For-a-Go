package main

// import "fmt"

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Good morning")
	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i,card)
	// }
	cards.print()
}

func newCard() string {
	return "Hello"
}
