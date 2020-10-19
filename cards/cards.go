package cards

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	Spades = iota
	Hearts
	Clubs
	Diamonds
)

var SuitSymbols = map[int]string{
	Spades:   "S",
	Hearts:   "H",
	Diamonds: "D",
	Clubs:    "C",
}

var SuitValues = map[string]int {
	"S": Spades,
	"H": Hearts,
	"D": Diamonds,
	"C": Clubs,
}

const (
	Ace   = 1
	Ten   = 10
	Jack  = 11
	Queen = 12
	King  = 13
)

var PipSymbols = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}
var PipValues = map[string]int {
	"A": Ace,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": Ten,
	"J": Jack,
	"Q": Queen,
	"K": King,
}

type Card struct {
	Suit  int
	Value int
}

func (card *Card) String() string {
	pipSymbol := PipSymbols[card.Value-1]
	suitSymbol := SuitSymbols[card.Suit]
	return fmt.Sprintf("%s%s", pipSymbol, suitSymbol)
}

type Pile []*Card

func (pile *Pile) String() string {
	var b strings.Builder
	for _, card := range *pile {
		_, _ = fmt.Fprintf(&b, "%s", card)
	}
	return b.String()
}

type Deck struct {
	Pile [52]*Card
	I    int
}

func (deck *Deck) Init() {
	for suit := 0; suit < 4; suit++ {
		for val := Ace; val <= King; val++ {
			i := (val - 1) + suit*13
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
