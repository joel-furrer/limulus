package parser

import (
	"fmt"
)

// tree icons
const (
	V = "│"
	T = "├"
	C = "└"
	H = "─"
)

const (
	Tee    = T + H + H + " "
	Corner = C + H + H + " "
	Pipe   = V + "   "
	Space  = "    "
)

func printAst(n Node, prefix string, isLast bool) {
	connector := Corner
	if !isLast {
		connector = Tee
	}

	switch n.Kind() {

	case NodeAssignment:
		fmt.Println("Assign")
		assign := n.(*AssignmentNode)

		fmt.Println(prefix + Tee + "Name: " + assign.Name)

		// recursion
		printAst(assign.Value, prefix, true)

	case NodeCout:
		fmt.Println("Cout")
		cout := n.(*CoutNode)

		// recursion
		printAst(cout.Value, prefix, true)

	case NodeNumber:
		num := n.(*NumberNode)

		fmt.Println(prefix + connector + fmt.Sprintf("Num[%v]: %v", num.NumType.ToString(), num.Value))

	case NodeIdentifier:
		id := n.(*IdentifierNode)

		fmt.Println(prefix + connector + "Id: " + id.Name)

	case NodeBinOp:
		op := n.(*BinOpNode)

		fmt.Println(prefix + connector + fmt.Sprintf("BinOp (%s)", op.Operator.ToString()))

		// recursion
		printAst(op.Left, prefix+getPrefix(isLast), false)
		// recursion
		printAst(op.Right, prefix+getPrefix(isLast), true)

	}
}

func getPrefix(isLast bool) string {
	if isLast {
		return Space
	} else {
		return Pipe
	}
}

/*

print ast function


take in a node
find its type and print

Assignment has 2 children:
- Name
- Expression

Cout has 1 child:
- Expression



*/
