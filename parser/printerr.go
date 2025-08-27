package parser

import (
	"fmt"
	"limulus/err"
)

func PrintErr(e err.Err, i Instruction) {
	i.Print()

	for range e.Position {
		fmt.Print(" ")
	}
	fmt.Println("^")
	currLine := i[0].Line
	fmt.Printf("%d:%d: %s\n", currLine, e.Position, e.Error)
}
