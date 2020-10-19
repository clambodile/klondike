package io

import (
	"../cards"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFoundations(t *testing.T) {
	t.Run("Given a string representation of foundations, returns 4 card piles.", func(t *testing.T) {
		input := "FAS2SF2H3HF3C4CF4D5D"
		expected := [4]cards.Pile{
			{
				{
					Suit:  cards.Spades,
					Value: cards.Ace,
				},
				{
					Suit:  cards.Spades,
					Value: 2,
				},
			},
			{
				{
					Suit:  cards.Hearts,
					Value: 2,
				},
				{
					Suit:  cards.Hearts,
					Value: 3,
				},
			},
			{
				{
					Suit:  cards.Clubs,
					Value: 3,
				},
				{
					Suit:  cards.Clubs,
					Value: 4,
				},
			},
			{
				{
					Suit:  cards.Diamonds,
					Value: 4,
				},
				{
					Suit:  cards.Diamonds,
					Value: 5,
				},
			},
		}
		foundations, err := parseFoundations(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, foundations)
	})
}

func TestParsePile(t *testing.T) {
	t.Run("Parses a pile of cards.", func(t *testing.T) {
		input := "AS3CKD"
		expected := cards.Pile{
			{
				Suit:  cards.Spades,
				Value: cards.Ace,
			},
			{
				Suit:  cards.Clubs,
				Value: 3,
			},
			{
				Suit:  cards.Diamonds,
				Value: cards.King,
			},
		}
		pile, err := parsePile(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, pile)
	})
}

func TestParseCard(t *testing.T) {
	t.Run("Parses a single card.", func(t *testing.T) {
		input := "AS"
		expected := cards.Card{
			Suit:  cards.Spades,
			Value: cards.Ace,
		}
		card, err := parseCard(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, card)
	})
}
