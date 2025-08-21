package lexer

import (
	"strconv"
	"unicode"

	"limulus/tok"
)

func classifyToken(text string) tok.Type {
	switch text {
	case "let":
		return tok.LET
	case "cout":
		return tok.COUT
	case "=":
		return tok.ASSIGN
	case "+":
		return tok.PLUS
	default:
		// check if is a number
		if _, err := strconv.Atoi(text); err == nil {
			return tok.NUMBER
		}

		// check if is a valid identifier
		if isValidIdentifier(text) {
			return tok.IDENTIFIER
		}

		return tok.NUMBER
	}
}

func isValidIdentifier(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
