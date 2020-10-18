package klondike

import "./cards"

type Column []*cards.Card

type Stack map[int][]*cards.Card

type Stock struct {
	Pile []*cards.Card
	I    int
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
		stock.Pile = append(stock.Pile[:stock.I], stock.Pile[stock.I + 1:]...)
		stock.I--
		if stock.I == -1 {
			stock.I = len(stock.Pile) - 1
		}
	}()
	return stock.Current()
}

type GameState struct {
	Foundations [4]Stack
	Tableau [7]Column
	Stock Stock
}

func (gameState *GameState) Init() *GameState {
	deck := &cards.Deck{}
	deck.Init()
	deck.Shuffle()
	gameState.Tableau = [7]Column{}
	for i := 0 ; i < 7; i++ {
		for j := 0; j < i + 1; j++ {
			gameState.Tableau[i] = append(gameState.Tableau[i], deck.Draw())
		}
	}
	for ; deck.I < len(deck.Pile); {
		gameState.Stock.Pile = append(gameState.Stock.Pile, deck.Draw())
	}
	return &GameState{}
}
