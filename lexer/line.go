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
				tokens = append(tokens, tok.Token{
					Text:     text,
					Position: start,
					Line:     l.LineNo,
					Type:     classifyToken(text),
				})
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
		tokens = append(tokens, tok.Token{
			Text:     text,
			Position: start,
			Line:     l.LineNo,
			Type:     classifyToken(text),
		})
	}

	return tokens
}
