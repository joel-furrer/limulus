package lexer

import (
	"unicode"

	"limulus/tok"
)

type Line struct {
	Text   string
	LineNo int
}

func (l Line) Tokens() []tok.Token {
	var tokens []tok.Token
	var inToken bool
	var start int
	var current []rune

	runes := []rune(string(l.Text))

	for i, r := range runes {
		if unicode.IsSpace(r) {
			if inToken {
				text := string(current)
				l.addToken(&tokens, text, start)
				current = nil
				inToken = false
			}
		} else {
			if !inToken {
				inToken = true
				start = i
			}
			current = append(current, r)
		}
	}

	if inToken {
		text := string(current)
		l.addToken(&tokens, text, start)
	}

	return tokens
}

func (l Line) addToken(tokens *[]tok.Token, text string, start int) {
	tokType := classifyToken(text)

	var numType tok.NumType
	if tokType == tok.NUMBER || tokType == tok.UNKNOWN {
		numType = classifyNumber(text)
		if numType != tok.NUM_UNKNOWN {
			tokType = tok.NUMBER
		}
	} else {
		numType = tok.NUM_UNKNOWN
	}

	*tokens = append(*tokens, tok.Token{
		Text:       text,
		Position:   start,
		Line:       l.LineNo,
		Type:       tokType,
		NumType:    numType,
	})
}
