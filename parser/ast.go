package parser

import (
	//"limulus/lexer"
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
