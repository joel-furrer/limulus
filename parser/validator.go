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
