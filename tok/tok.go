package tok

type Type int

const (
	UNKNOWN Type = iota
	
	IDENTIFIER

	NUMBER
	
	LET
	COUT
	
	ASSIGN
	DOT
	
	PLUS
	MIN
	MUL
	DIV
	
	LPAREN
	RPAREN
)

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

type Token struct {
	Text     string
	Position int
	Line     int
	Type     Type
	NumType  NumType
}
