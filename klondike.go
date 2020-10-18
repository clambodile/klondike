package klondike

import "./cards"

type Column []cards.Card

type Stack map[int][]cards.Card

type DrawPile struct {
	Pile []cards.Card
	I    int
}

type GameState struct {
	Stacks [4]Stack
	Columns [7]Column
	DrawPile DrawPile
}

func (gameState *GameState) Init() *GameState {
	deck := &cards.Deck{}
	//index of next card to draw
	dI := 0
	deck.Init()
	deck.Shuffle()
	var columns [7]Column
	for i := 0 ; i < 7; i++ {
		for j := 0; j <= i + 1; j++ {
			columns[i] = append(columns[i], deck.Pile[dI])
			dI++
		}
	}
	return &GameState{}
}
