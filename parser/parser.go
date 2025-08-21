package parser

import (
	//"limulus/lexer"
)

func Parse(instructions Instructions) (Program, error) {
	var prog Program

	_ = instructions.Validate()

	/*

	for _, instr := range instructions {
		node, err := instr.Ast()
	}

	*/


	return prog, nil
}
