package parser

import (
	"fmt"

	"limulus/tok"
	"limulus/err"
)

func ( instr Instructions ) Validate() bool {
	for _, i := range  instr {
		
		firstTok := i[0]
		if v, ok := statementValidators[firstTok.Type]; ok {
			err := v(i)
			if err.Error != nil {
				PrintErr(err, i)
				return false
			}
		} else {
			fmt.Printf("unknown token: %v\n", firstTok.Text)
			return false
		}
	}
	return true
}

type Validator func(i Instruction) err.Err

var statementValidators = map[tok.Type]Validator{
	tok.LET:  validateLet,
	tok.COUT: validateCout,
}

func validateLet(i Instruction) err.Err {

	if len(i) < 4 {
		lastTok := i[len(i) -1]
		pos := lastTok.Position + len(lastTok.Text) + 1
		return ErrNotEnoughArgs(tok.LET, pos)
	}

	if i[2].Type != tok.ASSIGN {
		pos := i[2].Position
		return ErrExpectedTokenAfter(tok.LET, tok.ASSIGN, pos)
	}

	tokErr := validateExpression(i[3:])

	if tokErr.Error != nil {
		// if validateExpression cant find a error position it uses the position where the expression starts
		if tokErr.Position == 0 {
			tokErr.Position += i[2].Position + len(i[2].Text) + 1
		}
		return tokErr
	}

	return err.Err{}
}

func validateCout(i Instruction) err.Err {

	if len(i) < 2 {
		pos := i[0].Position + len(i[0].Text) + 1
		return ErrNotEnoughArgs(tok.COUT, pos)
	}

	tokErr := validateExpression(i[1:])

	if tokErr.Error != nil {
		// if validateExpression cant find a error position it uses the position where the expression starts
		if tokErr.Position == 0 {
			tokErr.Position += i[0].Position + len(i[0].Text) + 1
		}
		return tokErr
	}

	return err.Err{}
}
