package parser

import (
	"fmt"

	//"limulus/lexer"
)

func Parse(instructions Instructions) (Program, error) {
	var prog Program

	if ok := instructions.Validate(); !ok {
		return prog, fmt.Errorf("validation failed")
	}

	for _, i := range instructions {
		//node, err := Instruction(i).Ast()
		node, _ := Instruction(i).Ast()

		fmt.Printf("%#v\n", node)
	}

	return prog, nil
}
