package parser

import (
	"fmt"
)

func Parse(instructions Instructions) (Program, error) {
	var prog Program

	if ok := instructions.Validate(); !ok {
		return prog, fmt.Errorf("validation failed")
	}

	for _, i := range instructions {
		node, _ := Instruction(i).Ast()

		printAst(node, "", true)
		fmt.Println()
	}

	return prog, nil
}
