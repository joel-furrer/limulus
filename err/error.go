package err

import "fmt"

type Err struct {
	Error    error
	Position int
}

func New(error string, pos int) Err {
	return Err{Error: fmt.Errorf(error), Position: pos}
}
