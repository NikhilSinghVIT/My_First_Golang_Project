package main

import "fmt"

func main() {

	cards := newDeck()
	fmt.Println("Cards before shuffling")
	cards.print()

	cards.shuffle()
	fmt.Println("cards after shuffling")

	cards.print()
}
