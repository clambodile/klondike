package klondike

import (
	"./cards"
	"fmt"
	"strings"
)

type Pile []*cards.Card

type Stock struct {
	Pile Pile
	I    int
}

func (pile *Pile) String() string {
	var b strings.Builder
	for _, card := range *pile {
		_, _ = fmt.Fprintf(&b, "%s", card)
	}
	return b.String()
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

type GameState struct {
	Foundations [4]Pile
	Tableau     [7]Pile
	Stock       Stock
}

func (gameState *GameState) Init() *GameState {
	deck := &cards.Deck{}
	deck.Init()
	deck.Shuffle()
	gameState.Tableau = [7]Pile{}
	for i := 0; i < 7; i++ {
		for j := 0; j < i+1; j++ {
			gameState.Tableau[i] = append(gameState.Tableau[i], deck.Draw())
		}
	}
	for deck.I < len(deck.Pile) {
		gameState.Stock.Pile = append(gameState.Stock.Pile, deck.Draw())
	}
	return &GameState{}
}

func (gameState *GameState) String() string {
	var b strings.Builder
	for _, foundation := range gameState.Foundations {
		_, _ = fmt.Fprintf(&b, "F%s", foundation.String())
	}
	for _, column := range gameState.Tableau {
		_, _ = fmt.Fprintf(&b, "T%s", column.String())
	}
	_, _ = fmt.Fprintf(&b, "S%s", gameState.Stock.String())
	return b.String()
}

//func ParseGameState(str string) *GameState {
//	pipPattern := "[ATJQK2-9]"
//	suitPattern := "[CDHS]"
//	cardPattern := fmt.Sprintf("(?:%s%s)", pipPattern, suitPattern)
//	foundationPattern := fmt.Sprintf("(?:F%s*)", cardPattern)
//	foundationsPattern := fmt.Sprintf("(%s{4})", foundationPattern)
//	columnPattern := fmt.Sprintf("(?:T%s*)", cardPattern)
//	tableauPattern := fmt.Sprintf("(%s{7})", columnPattern)
//	stockPattern := fmt.Sprintf("(S%s*)")
//	pattern := fmt.Sprintf("%s%s%s", foundationsPattern, tableauPattern, stockPattern)
//	readerRegexp := regexp.MustCompile(pattern)
//	return &GameState{}
//}
