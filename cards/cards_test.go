package cards

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShow(t *testing.T) {
	t.Run("Returns the correct string for any given card shown.", func(t *testing.T) {
		card1 := Card{Suit: spades, Value: ace}
		card2 := Card{Suit: hearts, Value: jack}
		card3 := Card{Suit: diamonds, Value: queen}
		card4 := Card{Suit: clubs, Value: king}
		card5 := Card{Suit: spades, Value: 2}

		assert.Equal(t, "AS", card1.Show())
		assert.Equal(t, "JH", card2.Show())
		assert.Equal(t, "QD", card3.Show())
		assert.Equal(t, "KC", card4.Show())
		assert.Equal(t, "2S", card5.Show())
	})
}

func TestInit(t *testing.T) {
	t.Run("Fills an empty Deck.", func(t *testing.T) {
		deck := &Deck{}
		assert.Empty(t, deck)
		deck.Init()
		assert.Len(t, deck.Pile, 52)
		for i := 0; i < len(deck.Pile); i++ {
			assert.NotNil(t, deck.Pile[i])
		}
	})
}

func TestShuffle(t *testing.T) {
	t.Run("Changes the order of the deck.", func(t *testing.T) {
		deck := &Deck{}
		deck.Init()
		var firstRun [52]string
		for i := 0; i < len(deck.Pile); i++ {
			firstRun[i] = deck.Pile[i].Show()
		}
		deck.Shuffle()
		countSame := 0
		for i := 0; i < len(deck.Pile); i++ {
			if deck.Pile[i].Show() == firstRun[i] {
				countSame++
			}
		}
		assert.Less(t, countSame, 5)
	})
}

func TestDraw(t *testing.T) {
	t.Run("Returns successive cards when called successively.", func(t *testing.T) {
		deck := &Deck{}
		deck.Init()
		assert.Equal(t, "AS", deck.Draw().Show())
		assert.Equal(t, "2S", deck.Draw().Show())
		assert.Equal(t, "3S", deck.Draw().Show())
	})
}
