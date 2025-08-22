package parser

import (
	"fmt"
	"limulus/tok"
)

func ( instr Instructions ) Validate() bool {
	for _, i := range  instr {
		
		firstTok := i[0]
		if v, ok := statementValidators[firstTok.Type]; ok {
			err := v(i)
			if err.Error != nil {
				err.Print(i)
				return false
			}
		} else {
			fmt.Printf("no validator for token: %v\n", firstTok.Text)
			return false
		}
	}
	return true
}

type Validator func(i Instruction) TokErr

var statementValidators = map[tok.Type]Validator{
	tok.LET:  validateLet,
	tok.COUT: validateCout,
}

func validateLet(i Instruction) TokErr {

	if len(i) < 4 {
		lastTok := i[len(i) -1]
		pos := lastTok.Position + len(lastTok.Text) + 1
		return NewTokErr("not enough arguments to call 'let'", pos)
	}

	if i[2].Type != tok.ASSIGN {
		pos := i[2].Position
		return NewTokErr("missing '=' after let statement", pos)
	}

	tokErr := validateExpression(i[3:])

	if tokErr.Error != nil {
		// if validateExpression cant find a error position it uses the position where the expression starts
		if tokErr.Position == 0 {
			tokErr.Position += i[2].Position + len(i[2].Text) + 1
		}
		return tokErr
	}

	return TokErr{}
}

func validateCout(i Instruction) TokErr {

	if len(i) < 2 {
		pos := i[0].Position + len(i[0].Text) + 1
		return NewTokErr("not enough arguments to call 'cout'", pos)
	}

	tokErr := validateExpression(i[1:])

	if tokErr.Error != nil {
		// if validateExpression cant find a error position it uses the position where the expression starts
		if tokErr.Position == 0 {
			tokErr.Position += i[0].Position + len(i[0].Text) + 1
		}
		return tokErr
	}

	return TokErr{}
}

type Expression Instruction

func validateExpression(i Instruction) TokErr {
	var e Expression = Expression(i)

	fmt.Println(e)

	// cant be nil
	if e == nil {
		return NewTokErr("missing expression", 0)
	}

	// must only contain valid tokens
	if pos, ok := e.ValidateTokens(); !ok  {
		return NewTokErr("invalid token", pos)
	}

	return TokErr{}
}

func (e Expression) ValidateTokens() (int, bool) {
	
	validTokens := []tok.Type{
		tok.IDENTIFIER,
		tok.NUMBER,
		tok.PLUS,
		tok.MIN,
		tok.MUL,
		tok.DIV,
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
