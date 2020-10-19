package game

import (
	"../cards"
	"fmt"
	"strings"
)

//The Stock is the deck of cards which may be viewed on or three
//at a time, and placed onto a Foundation or a Tableau.
type Stock struct {
	Pile cards.Pile
	I    int
}

func (stock *Stock) String() string {
	return stock.Pile.String()
}

func (stock *Stock) Current() *cards.Card {
	return stock.Pile[stock.I]
}

func (stock *Stock) Next() {
	stock.I++
	if stock.I == len(stock.Pile) {
		stock.I = 0
	}
}

func (stock *Stock) Draw() *cards.Card {
	defer func() {
		stock.Pile = append(stock.Pile[:stock.I], stock.Pile[stock.I+1:]...)
		stock.I--
		if stock.I == -1 {
			stock.I = len(stock.Pile) - 1
		}
	}()
	return stock.Current()
}

type State struct {
	Foundations [4]cards.Pile
	Tableau     [7]cards.Pile
	Stock       Stock
}

func (state *State) Init() *State {
	deck := &cards.Deck{}
	deck.Init()
	deck.Shuffle()
	state.Tableau = [7]cards.Pile{}
	for i := 0; i < 7; i++ {
		for j := 0; j < i+1; j++ {
			state.Tableau[i] = append(state.Tableau[i], deck.Draw())
		}
	}
	for deck.I < len(deck.Pile) {
		state.Stock.Pile = append(state.Stock.Pile, deck.Draw())
	}
	return &State{}
}

func (state *State) String() string {
	var b strings.Builder
	for _, foundation := range state.Foundations {
		_, _ = fmt.Fprintf(&b, "F%s", foundation.String())
	}
	for _, column := range state.Tableau {
		_, _ = fmt.Fprintf(&b, "T%s", column.String())
	}
	_, _ = fmt.Fprintf(&b, "S%s", state.Stock.String())
	return b.String()
}
