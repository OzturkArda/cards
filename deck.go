package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck [] string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three","Four","Five", "Six", "Seven", "Eight", "Nine", "Ten","Jack","Queen", "King"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print () {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) shuffle () {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}


func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte (d.toString()), 0666)
}

func deal( d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}


func (d deck) toString() string {
	return strings.Join([]string (d),",")
}

func newDeckFromFile(fileName string ) deck  {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil{
		fmt.Println("ERROR : ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs),",")
}