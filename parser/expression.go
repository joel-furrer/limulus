package parser

import (

	// temp
	"fmt"
	"limulus/tok"
)

type Expression Instruction

func validateExpression(i Instruction) TokErr {
	var e Expression = Expression(i)

	// cant be nil
	if e == nil {
		return NewTokErr("missing expression", 0)
	}

	// must only contain valid tokens
	if pos, ok := e.ValidateTokens(); !ok  {
		return NewTokErr("invalid token", pos)
	}

	// check parantheses
	tokErr := e.ValidateParantheses()
	if tokErr.Error != nil {
		return tokErr
	}

	// check for correct usage of binary operators
	tokErr = e.ValidateOperators()
	if tokErr.Error != nil {
		return tokErr
	}

	return TokErr{}
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

func (e Expression) ValidateParantheses() TokErr {
	var stack []tok.Token

	for i, t := range e {
		switch t.Type{

		case tok.LPAREN:
			if i == len(e) -1 {
				pos := t.Position + len(t.Text) -1
				return NewTokErr("cannot use '(' at the end of an expression", pos)
			}
			stack = append(stack, t)

		case tok.RPAREN:
			if i == 0 {
				return NewTokErr("cannot use ')' at the start of an expression", 0)
			}
			if len(stack) < 1 {
				return NewTokErr("missing '('", t.Position)
			}
			stack = stack[:len(stack) -1]
		}
	}

	if len(stack) != 0 {
		last := stack[len(stack) -1]
		return NewTokErr("missing ')'", last.Position)
	}

	return TokErr{}
}

func (e Expression) ValidateOperators() TokErr {

	var lastTok tok.Token
	for i, t := range e {
		if t.Type == tok.OP {
			fmt.Println("OP")
			if i == 0 {
				pos := t.Position
				return NewTokErr(fmt.Sprintf("cannot use '%s' at the start of an expression", t.Text), pos)
			}

			if i == len(e) -1 {
				pos := t.Position
				return NewTokErr(fmt.Sprintf("cannot use '%s' at the end of an expression", t.Text), pos)
			}
			
			if lastTok.Type == tok.OP {
				pos := t.Position
				return NewTokErr("missing identifier for operator", pos)
			}

		}
			
		lastTok = t
	}

	// split expressions into sub-exressions

	return TokErr{}
}


