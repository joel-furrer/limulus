package parser

import (
	"fmt"

	"limulus/tok"
)

type Program struct {
	Instructions Instructions
	Name         string
	AST          []Node
}

type Instructions [][]tok.Token

type Instruction []tok.Token

func (i Instruction) Print() {
	posProg := 0
	tokIndex := 0
	lastTok := i[len(i)-1]
	lastPos := lastTok.Position + len(lastTok.Text)

	for range lastPos {
		if posProg == i[tokIndex].Position {
			fmt.Print(i[tokIndex].Text)
			posProg += len(i[tokIndex].Text)
			if tokIndex < len(i)-1 {
				tokIndex++
			}
		} else {
			fmt.Print(" ")
			posProg++
		}
	}

	fmt.Println()
}
