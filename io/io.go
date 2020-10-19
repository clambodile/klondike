package io

import (
	"../cards"
	"fmt"
	"regexp"
	"strings"
)

var pipPattern = "[ATJQK2-9]"
var suitPattern = "[CDHS]"
var cardPattern = fmt.Sprintf("(?:%s%s)", pipPattern, suitPattern)

func parseFoundations(str string) ([4]cards.Pile, error) {
	foundations := [4]cards.Pile{}
	foundationPattern := fmt.Sprintf("F(%s*)", cardPattern)
	foundationsPattern := strings.Repeat(foundationPattern, 4)
	ok, err := regexp.Match(foundationsPattern, []byte(str))
	if err != nil {
		return foundations, fmt.Errorf("error matching str\n%s", str)
	}
	if !ok {
		return foundations, fmt.Errorf("input does not match foundations pattern\n%s", str)
	}
	re := regexp.MustCompile(foundationPattern)
	foundationStrs := re.FindAllString(str, -1)
	for i, foundationStr := range foundationStrs {
		foundation, err := parsePile(foundationStr)
		if err != nil {
			return foundations, err
		}
		foundations[i] = foundation
	}
	return foundations, nil
}

func parsePile(str string) (cards.Pile, error) {
	re := regexp.MustCompile(cardPattern)
	matches := re.FindAllString(str, -1)
	pile := cards.Pile{}
	for _, match := range matches {
		card, err := parseCard(match)
		if err != nil {
			return pile, err
		}
		pile = append(pile, &card)
	}
	return pile, nil
}

func parseCard(str string) (cards.Card, error) {
	pattern := fmt.Sprintf("(%s)(%s)", pipPattern, suitPattern)
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(str)[1:]
	pipStr := matches[0]
	pip := cards.PipValues[pipStr]
	suitStr := matches[1]
	suit := cards.SuitValues[suitStr]
	return cards.Card{Suit: suit, Value: pip}, nil
}

//func ParseGameState(str string) *game.State {
//	cardPattern := fmt.Sprintf("(?:%s%s)", pipPattern, suitPattern)
//	foundationPattern := fmt.Sprintf("(?:F%s*)", cardPattern)
//	foundationsPattern := fmt.Sprintf("(%s{4})", foundationPattern)
//	columnPattern := fmt.Sprintf("(?:T%s*)", cardPattern)
//	tableauPattern := fmt.Sprintf("(%s{7})", columnPattern)
//	stockPattern := fmt.Sprintf("(S%s*)")
//	pattern := fmt.Sprintf("%s%s%s", foundationsPattern, tableauPattern, stockPattern)
//	readerRegexp := regexp.MustCompile(pattern)
//	return &game.State{}
//}
