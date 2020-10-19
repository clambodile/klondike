package game

import (
	"../cards"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestStock_String(t *testing.T) {
	t.Run("Returns a string version of the pile", func(t *testing.T) {
		stock := Stock{
			Pile: cards.Pile{
				{Suit: 0, Value: 7},
				{Suit: 1, Value: 10},
				{Suit: 2, Value: 11},
				{Suit: 3, Value: 12},
			},
			I: 0,
		}
		assert.Equal(t, "7STHJCQD", stock.String())
	})
}

func TestStock_Current(t *testing.T) {
	t.Run("Shows the current available card in the stock", func(t *testing.T) {
		deck := &cards.Deck{}
		deck.Init()
		stock := Stock{
			Pile: deck.Pile[:24],
			I:    0,
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
			I:    0,
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
			I:    0,
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
			I:    0,
		}
		card := stock.Draw()
		assert.Equal(t, "AS", card.String())
	})
	t.Run("Removes the drawn card from the stock.", func(t *testing.T) {
		deck := &cards.Deck{}
		deck.Init()
		stock := Stock{
			Pile: deck.Pile[:24],
			I:    0,
		}
		drawn := stock.Draw()
		assert.Len(t, stock.Pile, 23)
		for _, card := range stock.Pile {
			assert.NotEqual(t, drawn.String(), card.String())
		}
	})
}

func TestGameState_Init(t *testing.T) {
	state := &State{}
	state.Init()

	t.Run("Columns of tableau consist of 1, 2, 3, 4, 5, 6, 7 cards.", func(t *testing.T) {
		for i := 0; i < 7; i++ {
			assert.Len(t, state.Tableau[i], i+1)
		}
	})
	t.Run("Stock starts with 24 cards.", func(t *testing.T) {
		assert.Len(t, state.Stock.Pile, 24)
	})
}

func TestGameState_String(t *testing.T) {
	t.Run("Creates a string of the game state.", func(t *testing.T) {
		state := &State{}
		state.Init()
		startsWithFoundations := regexp.MustCompile(`FFFF.*`)
		has7Tableaus := regexp.MustCompile(`FFFF(?:T[AJKQTCDHS2-9]+){7}`)
		endsWithStock := regexp.MustCompile(`.*S[AJKQTCDHS2-9]{48}`)
		assert.Regexp(t, startsWithFoundations, state.String())
		assert.Regexp(t, has7Tableaus, state.String())
		assert.Regexp(t, endsWithStock, state.String())
	})
}

func TestParseGameState(t *testing.T) {

}
