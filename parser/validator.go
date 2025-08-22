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
	var tokErr TokErr

	if len(i) < 4 {
		tokErr.Error = fmt.Errorf("not enough arguments to call 'let'")
		lastTok := i[len(i) -1]
		tokErr.Position = lastTok.Position + len(lastTok.Text) + 1
		return tokErr
	}

	if i[2].Type != tok.ASSIGN {
		tokErr.Error = fmt.Errorf("missing '=' after let statement")
		tokErr.Position = i[2].Position
		return tokErr
	}

	tokErr = validateExpression(i[3:])

	return tokErr
}

func validateCout(i Instruction) TokErr {
	var tokErr TokErr

	if len(i) < 2 {
		tokErr.Error = fmt.Errorf("not enough arguments to call 'cout'")
		tokErr.Position = i[0].Position + len(i[0].Text) + 1
		return tokErr
	}

	return tokErr
}

func validateExpression(i Instruction) TokErr {
	var tokErr TokErr

	return tokErr
}
