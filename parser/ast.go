package parser

import (
	"fmt"

	"limulus/tok"
)

// AST
type NodeKind int

const (
	NodeAssignment NodeKind = iota
	NodeCout
	NodeBinOp
	NodeNumber
	NodeIdentifier
)

type Node interface {
	Kind() NodeKind
}

// AST Nodes

// ----- Assignment -----
type AssignmentNode struct {
	Name  string
	Value Node
}

func (a *AssignmentNode) Kind() NodeKind { return NodeAssignment }

func AsAssignment(n Node) *AssignmentNode {
	if n.Kind() != NodeAssignment {
		panic("not an AssignmentNode")
	}
	return n.(*AssignmentNode)
}

// ----- Cout -----
type CoutNode struct {
	Value Node
}

func (c *CoutNode) Kind() NodeKind { return NodeCout }

func AsCout(n Node) *CoutNode {
	if n.Kind() != NodeCout {
		panic("not an CoutNode")
	}
	return n.(*CoutNode)
}

// ----- BinOp -----
type BinOpNode struct {
	Left     Node
	Operator tok.BinOpType
	Right    Node
}

func (b *BinOpNode) Kind() NodeKind { return NodeBinOp }

func AsBinOp(n Node) *BinOpNode {
	if n.Kind() != NodeBinOp {
		panic("not an BinOpNode")
	}
	return n.(*BinOpNode)
}

// ----- Number -----
type NumberNode struct {
	Value   int
	NumType tok.NumType
}

func (n *NumberNode) Kind() NodeKind { return NodeNumber }

func AsNumber(n Node) *NumberNode {
	if n.Kind() != NodeNumber {
		panic("not an NumberNode")
	}
	return n.(*NumberNode)
}

// ----- Identifier -----
type IdentifierNode struct {
	Name string
}

func (i *IdentifierNode) Kind() NodeKind { return NodeIdentifier }

func AsIdentifier(n Node) *IdentifierNode {
	if n.Kind() != NodeIdentifier {
		panic("not an IdentifierNode")
	}
	return n.(*IdentifierNode)
}

// AST
func (i Instruction) Ast() (Node, error) {

	switch i[0].Type {
	case tok.LET:
		return i.AssignmentAst(), nil
	case tok.COUT:
		return i.CoutAst(), nil
	}

	return nil, nil
}

func (i Instruction) AssignmentAst() Node {
	pos := 3
	expr := parseExpression(i, &pos)
	return &AssignmentNode{Name: i[1].Text, Value: expr}
}

func (i Instruction) CoutAst() Node {
	pos := 1
	expr := parseExpression(i, &pos)
	//return &CoutNode{Name: i[1].Text, Value: expr}
	return &CoutNode{Value: expr}
}

func parseExpression(tokens Instruction, pos *int) Node {
	left := parseTerm(tokens, pos)

	for *pos < len(tokens) {
		t := tokens[*pos]

		// only stay if current token is operator
		if t.Type != tok.OP {
			break
		}

		// only stay if current token is - or +
		if t.BinOpType != tok.PLUS && t.BinOpType != tok.MIN {
			break
		}

		*pos++
		right := parseTerm(tokens, pos)
		left = &BinOpNode{Left: left, Operator: t.BinOpType, Right: right}
	}

	return left
}

func parseTerm(tokens Instruction, pos *int) Node {
	left := parseFactor(tokens, pos)

	for *pos < len(tokens) {
		t := tokens[*pos]

		// only stay if current token is operator
		if t.Type != tok.OP {
			break
		}

		// only stay if current token is * or /
		if t.BinOpType != tok.MUL && t.BinOpType != tok.DIV {
			break
		}

		*pos++
		right := parseFactor(tokens, pos)
		left = &BinOpNode{Left: left, Operator: t.BinOpType, Right: right}
	}

	return left
}

func parseFactor(tokens Instruction, pos *int) Node {
	t := tokens[*pos]
	*pos++

	switch t.Type {
	case tok.NUMBER:
		var val int
		fmt.Sscanf(t.Text, "%d", &val)
		return &NumberNode{Value: val, NumType: t.NumType}

	case tok.IDENTIFIER:
		return &IdentifierNode{Name: t.Text}
	case tok.LPAREN:
		node := parseExpression(tokens, pos)

		// this *pos++ skips the ')', where pos is still pointing to from parseExpression
		*pos++
		return node
	}

	return nil
}
