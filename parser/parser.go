package parser

import (
	"fmt"
)

type Options struct {
	DumpAST      bool
	DumpASTTyped bool
}

var DumpASTTyped bool

func Parse(instr Instructions, opts Options) (Program, error) {
	var prog Program

	if ok := instr.Validate(); !ok {
		return prog, fmt.Errorf("validation failed")
	}

	for _, i := range instr {
		node, _ := Instruction(i).Ast()

		analyzeNode(node)

		DumpASTTyped = opts.DumpASTTyped

		if opts.DumpAST || opts.DumpASTTyped {
			printAst(node, "", true)
		}
	}

	return prog, nil
}
