package klondike

import (
	"./cards"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStock_Current(t *testing.T) {
	t.Run("Shows the current available card in the stock", func(t *testing.T) {
		deck := &cards.Deck{}
		deck.Init()
		stock := Stock{
			Pile: deck.Pile[:24],
			I: 0,
		}
		assert.Equal(t, "AS", stock.Current().String())
	})
}

func TestStock_Next(t *testing.T) {
	t.Run("Cycles to the next card in the stock", func(t *testing.T) {
		deck := &cards.Deck{}
		deck.Init()
		stock := Stock{
			Pile: deck.Pile[:24],
			I: 0,
		}
		assert.Equal(t, "AS", stock.Current().String())
		stock.Next()
		assert.Equal(t, "2S", stock.Current().String())
	})
	t.Run("Returns to beginning after reaching the end.", func(t *testing.T) {
		deck := &cards.Deck{}
		deck.Init()
		stock := Stock{
			Pile: deck.Pile[:],
			I: 0,
		}
		assert.Equal(t, "AS", stock.Current().String())
		for i := 0; i < 51; i++ {
			stock.Next()
		}
		assert.Equal(t, "KD", stock.Current().String())
		stock.Next()
		assert.Equal(t, "AS", stock.Current().String())
	})
}

func TestStock_Draw(t *testing.T) {
	t.Run("Returns the current card", func(t *testing.T) {
		deck := &cards.Deck{}
		deck.Init()
		stock := Stock{
			Pile: deck.Pile[:24],
			I: 0,
		}
		card := stock.Draw()
		assert.Equal(t, "AS", card.String())
	})
	t.Run("Removes the drawn card from the stock.", func(t *testing.T) {
		deck := &cards.Deck{}
		deck.Init()
		stock := Stock{
			Pile: deck.Pile[:24],
			I: 0,
		}
		drawn := stock.Draw()
		assert.Len(t, stock.Pile, 23)
		for _, card := range stock.Pile {
			assert.NotEqual(t, drawn.String(), card.String())
		}
	})
}

func TestGameState_Init(t *testing.T) {
	gameState := &GameState{}
	gameState.Init()

	t.Run("Columns of tableau consist of 1, 2, 3, 4, 5, 6, 7 cards.", func(t *testing.T) {
		for i := 0; i < 7; i++ {
			assert.Len(t, gameState.Tableau[i], i + 1)
		}
	})
	t.Run("Stock starts with 24 cards.", func(t *testing.T) {
		assert.Len(t, gameState.Stock.Pile, 24)
	})
}