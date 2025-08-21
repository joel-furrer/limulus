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

	return tokErr
}

func validateCout(i Instruction) TokErr {
	var tokErr TokErr

	if len(i) < 2 {
		tokErr.Error = fmt.Errorf("not enough arguments to call 'cout'")
		tokErr.Position = i[0].Position + len(i[0].Text) + 1
		return tokErr
	}

	//tokErr = validateExpression(i[1:])

	return tokErr
}

/*
func validateExpression(i Instruction) TokErr {
	var tokErr TokErr
	tokErr.Line = i[0].Line
	var lastTok tok.Token

	for j, t := range i {
		switch t.Type {
		case tok.NUMBER:
			if lastTok.Type == tok.NUMBER || lastTok.Type == tok.IDENTIFIER {
				tokErr.Error = fmt.Errorf("missing operand")
				tokErr.Position = t.Position
				return tokErr
			}

		case tok.IDENTIFIER:
			if lastTok.Type == tok.NUMBER || lastTok.Type == tok.IDENTIFIER {
				tokErr.Error = fmt.Errorf("missing operand")
				tokErr.Position = t.Position
				return tokErr
			}

			_, exists := getVar(t.Text)
			if !exists {
				tokErr.Error = fmt.Errorf("undefined: %s", t.Text)
				tokErr.Position = t.Position
			}

		case tok.PLUS:
			if j == 0 {
				tokErr.Error = fmt.Errorf("no value before '+'")
				tokErr.Position = t.Position
				return tokErr
			}

			if j == len(i)-1 {
				tokErr.Error = fmt.Errorf("no value after '+'")
				tokErr.Position = t.Position + 2
				return tokErr
			}

			if lastTok.Type != tok.NUMBER && lastTok.Type != tok.IDENTIFIER {
				tokErr.Error = fmt.Errorf("cannot use '+' for '%s'", lastTok.Text)
				tokErr.Position = lastTok.Position
				return tokErr
			}

		default:
			tokErr.Error = fmt.Errorf("invalid token '%s' in expression", t.Text)
			tokErr.Position = t.Position
			return tokErr
		}
		lastTok = t
	}

	return tokErr
}
*/
