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

type CoutNode struct {
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

	switch i[0].Type{
	case tok.LET:
		return i.AssignmentAst(), nil
	case tok.COUT:
		return i.CoutAst(), nil
	}

	return nil, nil
}

func (i Instruction) AssignmentAst() AssignmentNode {
    pos := 3
    expr := parseExpression(i, &pos)
    return AssignmentNode{Name: i[1].Text, Value: expr}
}

func (i Instruction) CoutAst() CoutNode {
    pos := 3
    expr := parseExpression(i, &pos)
    return CoutNode{Name: i[1].Text, Value: expr}
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
		left = BinOpNode{Left: left, Operator: t.Text, Right: right}
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
        left = BinOpNode{Left: left, Operator: t.Text, Right: right}
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
        return NumberNode{Value: val}

    case tok.IDENTIFIER:
        return IdentifierNode{Name: t.Text}
	case tok.LPAREN:
		node := parseExpression(tokens, pos)
		
		// this *pos++ skips the ')', where pos is still pointing to from parseExpression
		*pos++
		return node
	}

	return nil
}
