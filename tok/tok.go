package tok

// tokens
type Type int

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
