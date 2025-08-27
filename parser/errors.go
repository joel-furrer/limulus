package parser

import (
	"fmt"

	"limulus/tok"
	"limulus/err"
)

// complete error messages
const (
	ErrMissingExpression = "missing expression"
	ErrInvalidToken = "invalid token"
)

// enums for context
const (
	AtStartOfExpression = "at the start of an expression"
	AtEndOfExpression = "at the end of an expression"
)

func ErrInvalidTokenUsage(t tok.Token, ctx string, pos int ) err.Err {
	return err.New(fmt.Sprintf("cannot use '%s' %s", t.Text, ctx), pos)
}

func ErrMissingToken(t tok.Type, pos int) err.Err {
	var text string
	switch t{
	case tok.LPAREN:
		text = "("
	case tok.RPAREN:
		text = ")"
	default:
		text = "unknown token"
	} 
	return err.New(fmt.Sprintf("missing '%s'", text), pos)
}

func ErrUnexpectedSequence(prev tok.Token, curr tok.Token, pos int ) err.Err {
	return err.New(fmt.Sprintf("unexpected '%s' after '%s'", curr.Text, prev.Text), pos)
}
