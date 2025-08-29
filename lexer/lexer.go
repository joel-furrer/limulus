package lexer

import "limulus/tok"

func Lex(source SourceFile) [][]tok.Token {
	var instructions [][]tok.Token

	for _, line := range source.Lines() {
		tokens := line.Tokens()
		instructions = append(instructions, tokens)
	}

	return instructions
}
