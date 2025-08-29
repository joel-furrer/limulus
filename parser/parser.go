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
	
	table := NewSymbolTable()
	
	DumpASTTyped = opts.DumpASTTyped

	for _, i := range instr {
		
		ast, _ := Instruction(i).Ast()

		if opts.DumpAST || opts.DumpASTTyped {
			printAst(ast, "", true)
		}

		_, err := analyzeNode(ast, table)
		if err != nil {
			return prog, err
		}
	}

	return prog, nil
}
