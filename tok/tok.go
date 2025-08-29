package tok

// tokens
type Type int

func (t Type) ToString() string {
	switch t {
	case IDENTIFIER:
		return "identifier"
	case NUMBER:
		return "number"
	case LET:
		return "let"
	case COUT:
		return "cout"
	case ASSIGN:
		return "="
	case OP:
		return "operator"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	default:
		return "unknown token"
	}
}

const (
	UNKNOWN Type = iota

	IDENTIFIER

	NUMBER

	LET
	COUT

	ASSIGN
	DOT

	OP

	LPAREN
	RPAREN
)

// numbers
type NumType int

const (
	NUM_UNKNOWN NumType = iota
	INT32
	INT64
	UINT32
	UINT64
	FLOAT32
	FLOAT64
)

// binary operators
type BinOpType int

/*
func ( b BinOpType ) ToString() string {
	switch b {
	case PLUS:
		return "+"
	case MIN:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	}
}
*/

const (
	OP_UNKNOWN BinOpType = iota
	PLUS
	MIN
	MUL
	DIV
)

type Token struct {
	Text      string
	Position  int
	Line      int
	Type      Type
	NumType   NumType
	BinOpType BinOpType
}
