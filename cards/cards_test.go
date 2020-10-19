package cards

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCard_String(t *testing.T) {
	t.Run("Returns the correct string for any given card shown.", func(t *testing.T) {
		card1 := Card{Suit: Spades, Value: ace}
		card2 := Card{Suit: Hearts, Value: Jack}
		card3 := Card{Suit: Diamonds, Value: Queen}
		card4 := Card{Suit: Clubs, Value: King}
		card5 := Card{Suit: Spades, Value: 2}

		assert.Equal(t, "AS", card1.String())
		assert.Equal(t, "JH", card2.String())
		assert.Equal(t, "QD", card3.String())
		assert.Equal(t, "KC", card4.String())
		assert.Equal(t, "2S", card5.String())
	})
}

func TestPile_String(t *testing.T) {
	t.Run("Returns a string version of the pile", func(t *testing.T) {
		pile := Pile{
			{Suit: 0, Value: 1},
			{Suit: 1, Value: 2},
			{Suit: 2, Value: 3},
			{Suit: 3, Value: 13},
		}
		assert.Equal(t, "AS2H3CKD", pile.String())
	})
}

func TestDeck_Init(t *testing.T) {
	t.Run("New decks are empty.", func(t *testing.T) {
		deck := &Deck{}
		assert.Empty(t, deck)
	})
	t.Run("Fills an empty Deck.", func(t *testing.T) {
		deck := &Deck{}
		deck.Init()
		assert.Len(t, deck.Pile, 52)
		for i := 0; i < len(deck.Pile); i++ {
			assert.NotNil(t, deck.Pile[i])
		}
	})
}

func TestDeck_Shuffle(t *testing.T) {
	t.Run("Changes the order of the deck.", func(t *testing.T) {
		deck := &Deck{}
		deck.Init()
		var firstRun [52]string
		for i := 0; i < len(deck.Pile); i++ {
			firstRun[i] = deck.Pile[i].String()
		}
		deck.Shuffle()
		countSame := 0
		for i := 0; i < len(deck.Pile); i++ {
			if deck.Pile[i].String() == firstRun[i] {
				countSame++
			}
		}
		assert.Less(t, countSame, 5)
	})
}

func TestDeck_Draw(t *testing.T) {
	t.Run("Returns successive cards when called successively.", func(t *testing.T) {
		deck := &Deck{}
		deck.Init()
		assert.Equal(t, "AS", deck.Draw().String())
		assert.Equal(t, "2S", deck.Draw().String())
		assert.Equal(t, "3S", deck.Draw().String())
	})
}
