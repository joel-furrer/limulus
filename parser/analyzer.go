package parser

import (
	"fmt"

	"limulus/tok"
)

type Symbol struct {
	Name string
	Type tok.NumType
}

type SymbolTable struct {
	Vars map[string]Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{Vars: make(map[string]Symbol)}
}

func analyzeNode(n Node, table *SymbolTable) (tok.NumType, error) {

	switch n.Kind() {

	case NodeAssignment:
		node := AsAssignment(n)
		valType, err := analyzeNode(node.Value, table)
		if err != nil {
			return tok.NUM_UNKNOWN, err
		}
		table.Vars[node.Name] = Symbol{Name: node.Name, Type: valType}
		return valType, nil

	case NodeCout:
		node := AsCout(n)
		return analyzeNode(node.Value, table)

	case NodeNumber:
		node := AsNumber(n)
		return node.NumType, nil

	case NodeBinOp:
		node := AsBinOp(n)
		lType, err := analyzeNode(node.Left, table)
		if err != nil {
			return tok.NUM_UNKNOWN, err
		}
		rType, err := analyzeNode(node.Right, table)
		if err != nil {
			return tok.NUM_UNKNOWN, err
		}

		lStr := lType.ToString()
		rStr := rType.ToString()

		lVal := valueToString(node.Left)
		rVal := valueToString(node.Right)

		if lType == rType {
			fmt.Printf("[TYPE CHECKER]: success: %s[%s] == %s[%s]\n", lVal, lStr, rVal, rStr)
			// implement type checker that promotes type if needed or returns error ( consider uint, int, float, negative nums )
			// important: uints are always set explict, therefore should not be promoted to a int
			return lType, nil
		} else {
			return tok.NUM_UNKNOWN, fmt.Errorf("type conflict: %s[%s] != %s[%s]\n", lVal, lStr, rVal, rStr)
		}

	case NodeIdentifier:
		node := AsIdentifier(n)
		sym, ok := table.Vars[node.Name]
		if !ok {
			fmt.Println("identifier does not exist:", node.Name)
		}
		return sym.Type, nil

	}

	return tok.NUM_UNKNOWN, nil
}

func valueToString(n Node) string {
	switch n.Kind() {
	case NodeNumber:
		num := AsNumber(n)
		return fmt.Sprintf("%v", num.Value)
	case NodeIdentifier:
		id := AsIdentifier(n)
		return id.Name
	case NodeBinOp:
		bin := AsBinOp(n)
		return fmt.Sprintf("(%s %v %s)",
			valueToString(bin.Left),
			bin.Operator.ToString(),
			valueToString(bin.Right))
	default:
		return "<?>"
	}
}
