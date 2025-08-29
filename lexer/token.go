package lexer

import (
	"strconv"
	"strings"
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

	// operands
	case "+", "-", "*", "/":
		return tok.OP

	// parentheses
	case "(":
		return tok.LPAREN
	case ")":
		return tok.RPAREN

	default:
		// check if is a valid identifier
		if isValidIdentifier(text) {
			return tok.IDENTIFIER
		}

		// check if is a number
		if _, err := strconv.Atoi(text); err == nil {
			return tok.NUMBER
		}

		return tok.UNKNOWN
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

func classifyNumber(text string) tok.NumType {

	text = strings.ToLower(text)

	hasSign := false
	if strings.HasPrefix(text, "+") || strings.HasPrefix(text, "-") {
		text = text[1:]
		hasSign = true
	}

	// check if is float
	if strings.Contains(text, ".") {
		if strings.HasSuffix(text, "u") {
			return tok.NUM_UNKNOWN
		}

		if _, err := strconv.ParseFloat(text, 64); err != nil {
			return tok.FLOAT64
		}
		return tok.NUM_UNKNOWN
	}

	// check if is unsigned
	unsigned := false
	if strings.HasSuffix(text, "u") {
		unsigned = true
		text = text[:len(text)-1]
	}

	if hasSign && unsigned {
		return tok.NUM_UNKNOWN
	}

	if _, err := strconv.ParseInt(text, 10, 32); err == nil {
		if unsigned {
			return tok.UINT32
		}
		return tok.INT32
	}

	return tok.NUM_UNKNOWN
}

func classifyOperator(text string) tok.BinOpType {
	switch text {
	case "+":
		return tok.PLUS
	case "-":
		return tok.MIN
	case "*":
		return tok.MUL
	case "/":
		return tok.DIV
	default:
		return tok.OP_UNKNOWN
	}
}
