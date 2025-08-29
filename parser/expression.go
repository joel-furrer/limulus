package parser

import (
	"limulus/err"
	"limulus/tok"
)

type Expression Instruction

func validateExpression(i Instruction) err.Err {
	var e Expression = Expression(i)

	if e == nil {
		return err.New(ErrMissingExpression, 0)
	}

	if pos, ok := e.ValidateTokens(); !ok {
		return err.New(ErrInvalidToken, pos)
	}

	tokErr := e.ValidateParantheses()
	if tokErr.Error != nil {
		return tokErr
	}

	tokErr = e.ValidateOperators()
	if tokErr.Error != nil {
		return tokErr
	}

	tokErr = e.ValidateSequence()
	if tokErr.Error != nil {
		return tokErr
	}

	return err.Err{}
}

func (e Expression) ValidateTokens() (int, bool) {

	validTokens := []tok.Type{
		tok.IDENTIFIER,
		tok.NUMBER,
		tok.OP,
		tok.LPAREN,
		tok.RPAREN,
	}

	tokMap := make(map[tok.Type]bool)
	for _, vt := range validTokens {
		tokMap[vt] = true
	}

	for _, t := range e {
		if !tokMap[t.Type] {
			return t.Position, false
		}
	}

	return 0, true
}

func (e Expression) ValidateParantheses() err.Err {
	var stack []tok.Token
	var lastTok tok.Token

	for i, t := range e {
		switch t.Type {

		case tok.LPAREN:
			if i == len(e)-1 {
				pos := t.Position + len(t.Text) - 1
				return ErrInvalidTokenUsage(t, AtEndOfExpression, pos)
			}

			if lastTok.Type == tok.RPAREN {
				return ErrExpectedTokenBetween(lastTok.Type, t.Type, tok.OP, lastTok.Position)
			}

			stack = append(stack, t)

		case tok.RPAREN:
			if i == 0 {
				return ErrInvalidTokenUsage(t, AtStartOfExpression, 0)
			}

			if len(stack) < 1 {
				return ErrMissingToken(tok.LPAREN, t.Position)
			}

			if lastTok.Type == tok.LPAREN {
				return err.New(ErrEmptyExpression, lastTok.Position)
			}

			stack = stack[:len(stack)-1]
		}

		lastTok = t

	}

	if len(stack) != 0 {
		last := stack[len(stack)-1]
		return ErrMissingToken(tok.RPAREN, last.Position)
	}

	return err.Err{}
}

func (e Expression) ValidateOperators() err.Err {

	var lastTok tok.Token
	for i, t := range e {
		if t.Type == tok.OP {
			if i == 0 {
				pos := t.Position
				return ErrInvalidTokenUsage(t, AtStartOfExpression, pos)
			}

			if i == len(e)-1 {
				pos := t.Position
				return ErrInvalidTokenUsage(t, AtEndOfExpression, pos)
			}

			if lastTok.Type == tok.OP {
				pos := t.Position
				return ErrUnexpectedSequence(lastTok, t, pos)
			}

			if lastTok.Type == tok.LPAREN {
				pos := t.Position
				return ErrUnexpectedSequence(lastTok, t, pos)
			}

		}

		lastTok = t
	}

	return err.Err{}
}

func (e Expression) ValidateSequence() err.Err {
	var lastTok tok.Token

	for _, t := range e {
		switch t.Type {

		case tok.NUMBER, tok.IDENTIFIER:
			if lastTok.Type == tok.NUMBER || lastTok.Type == tok.IDENTIFIER || lastTok.Type == tok.RPAREN {
				return ErrUnexpectedSequence(lastTok, t, t.Position)
			}

		case tok.LPAREN:
			if lastTok.Type == tok.NUMBER || lastTok.Type == tok.IDENTIFIER || lastTok.Type == tok.RPAREN {
				return ErrUnexpectedSequence(lastTok, t, t.Position)
			}

		case tok.RPAREN:
			if lastTok.Type == tok.OP {
				return ErrUnexpectedSequence(lastTok, t, t.Position)
			}

		case tok.OP:
			if lastTok.Type == tok.LPAREN {
				return ErrUnexpectedSequence(lastTok, t, t.Position)
			}

		}

		lastTok = t
	}

	return err.Err{}
}
