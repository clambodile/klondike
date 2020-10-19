package io

//import (
//	"../cards"
//	"../game"
//	"regexp"

//"fmt"
//"regexp"
//)

//var pipPattern = "[ATJQK2-9]"
//var suitPattern = "[CDHS]"

//func parseFoundations(str string) ([4]cards.Pile, error) {
//	foundationsPattern =
//}

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
