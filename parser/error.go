package parser

import "fmt"

type TokErr struct {
	Error    error
	Position int
}

func (te TokErr) Print(i Instruction) {
	i.Print()

	for range te.Position {
		fmt.Print(" ")
	}
	fmt.Println("^")
	currLine := i[0].Line
	fmt.Printf("%d:%d: %s\n", currLine, te.Position, te.Error)
}

func NewTokErr(error string, pos int) TokErr {
	return TokErr{Error: fmt.Errorf(error), Position: pos}
}
