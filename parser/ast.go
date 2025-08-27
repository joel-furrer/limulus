package parser

import (
	"fmt"

	"limulus/tok"
)

// AST
type Node interface{}

type BinOpNode struct {
	Left     Node
	Operator string
	Right    Node
}

type AssignmentNode struct {
	Name  string
	Value Node
}

type NumberNode struct {
	Value int
}

type IdentifierNode struct {
	Name string
}

func ( i Instruction ) Ast() ( Node, error ) {
	fmt.Println(i)

	switch i[0].Type{
	case tok.LET:
		return i.AssignmentAst(), nil
	}


	return nil, nil
}

func ( i Instruction ) AssignmentAst() AssignmentNode {

/*	
if expression is just 1 element, decide if number of identifier node
else: parse expression using a recursive descent parser
*/

	return AssignmentNode{Name: i[1].Text}
}
