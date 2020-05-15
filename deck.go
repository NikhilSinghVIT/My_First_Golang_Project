package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

//No need to use reciever for this function
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamond", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for i := range cardSuits {
		for j := range cardValues {
			cards = append(cards, cardValues[j]+" of "+cardSuits[i])
		}
	}
	return cards
}

//Reciever d is used in this function
func (d deck) print() {
	for i := range d {
		fmt.Println(d[i])
	}
}

//function for dealing

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

//helper function to join the deck into a comma separated string without using join

// func (d deck) toString() string {
// 	var str string = ""
// 	for i := range d {
// 		str += "," + d[i]
// 	}
// 	return str
// }

//helper function to join the deck into a comma separated string using the join function

func (d deck) toString() string {
	return (strings.Join(d, ","))
}

//function to save the deck as a comma separated string into a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//function to load a deck from a file containing the deck as comma separated string

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
