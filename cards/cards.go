package cards

import (
	"fmt"
	"math/rand"
)

const (
	spades = iota
	hearts
	clubs
	diamonds
)

const (
	ace   = 1
	jack  = 11
	queen = 12
	king  = 13
)

type Card struct {
	Suit  int
	Value int
}

func (card *Card) String() string {
	pipSymbols := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	pipSymbol := pipSymbols[card.Value - 1]
	var suitSymbols [4]string
	suitSymbols[spades] = "S"
	suitSymbols[hearts] = "H"
	suitSymbols[diamonds] = "D"
	suitSymbols[clubs] = "C"
	suitSymbol := suitSymbols[card.Suit]
	return fmt.Sprintf("%s%s", pipSymbol, suitSymbol)

}

type Deck struct {
	Pile [52]*Card
	I int
}
func (deck *Deck) Init() {
	for suit := 0; suit < 4; suit++ {
		for val := ace; val <= king; val++ {
			i := (val - 1) + suit * 13
			deck.Pile[i] = &Card{Suit: suit, Value: val}
		}
	}
	deck.I = 0
}
func (deck *Deck) Shuffle() {
	rand.Shuffle(52, func(i, j int) {
		deck.Pile[i], deck.Pile[j] = deck.Pile[j], deck.Pile[i]
	})
	deck.I = 0
}
func (deck *Deck) Draw() *Card {
	defer func() { deck.I++ }()
	return deck.Pile[deck.I]
}
