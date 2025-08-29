package parser

import (
	"fmt"
)

func printAst(n Node) {
	switch n.Kind() {
	case NodeAssignment:
		fmt.Println("assign node")
	case NodeCout:
		fmt.Println("cout node")
	}
}
