package tok

type Type int

const (
	UNKNOWN Type = iota
	NUMBER
	IDENTIFIER
	LET
	ASSIGN
	PLUS
	COUT
)

type Token struct {
	Text     string
	Position int
	Line     int
	Type     Type
}
