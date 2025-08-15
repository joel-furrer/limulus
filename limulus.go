package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// FileData
type SourceFile struct {
	FileName string
	Content  string
}

func (f *SourceFile) Lines() []Line {
	var lines []Line
	scanner := bufio.NewScanner(strings.NewReader(f.Content))
	lineNo := 1

	for scanner.Scan() {
		if scanner.Text() != "" {
			lines = append(lines, Line{
				Text:	scanner.Text(),
				LineNo: lineNo,
			})
		}
		lineNo++
	}

	return lines
}

// tokens
type tokenType int

const (
	TOK_UNKNOWN tokenType = iota
	TOK_NUMBER
	TOK_IDENTIFIER
	TOK_LET
	TOK_ASSIGN
	TOK_PLUS
	TOK_COUT
)

func classifyToken(text string) tokenType {
	switch text {
	case "let":
		return TOK_LET
	case "cout":
		return TOK_COUT
	case "=":
		return TOK_ASSIGN
	case "+":
		return TOK_PLUS
	default:
		// check if is a number
		if _, err := strconv.Atoi(text); err == nil {
			return TOK_NUMBER
		}

		// check if is a valid identifier
		if isValidIdentifier(text) {
			return TOK_IDENTIFIER
		}

		return TOK_NUMBER
	}
}

func isValidIdentifier(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

type Token struct {
	Text     string
	Position int
	Line     int
	Type     tokenType
}

type Line struct {
	Text string
	LineNo int
}

func (l Line) Tokens() []Token {
	var tokens []Token
	var inToken bool
	var start int
	var current []rune

	runes := []rune(string(l.Text))

	for i, r := range runes {
		if unicode.IsSpace(r) {
			if inToken {
				text := string(current)
				tokens = append(tokens, Token{
					Text:     text,
					Position: start,
					Line:     l.LineNo,
					Type:     classifyToken(text),
				})
				current = nil
				inToken = false
			}
		} else {
			if !inToken {
				inToken = true
				start = i
			}
			current = append(current, r)
		}
	}

	if inToken {
		text := string(current)
		tokens = append(tokens, Token{
			Text:     text,
			Position: start,
			Line:     l.LineNo,
			Type:     classifyToken(text),
		})
	}

	return tokens
}

func cout(value int) {
	fmt.Println(value)
}

// validation functions
func validate_let(i Instruction) error {
	if len(i) < 3 {
		return fmt.Errorf("not enough arguments to call 'let'")
	}
	return nil
}

func validate_cout(i Instruction, t Token) ErrPosition {
	var errPos ErrPosition
	if len(i) == 0 {
		errPos.Error = fmt.Errorf("missing argument for 'cout'")
		errPos.Line = t.Line
		errPos.Position = t.Position + len(t.Text)
	} //else if len(i) > 1 {
		//return fmt.Errorf("too many arguments to call 'cout'")
	//}

	return errPos 
}

type ErrPosition struct {
	Error error
	Line int
	Position int
}

func (i Instruction) Validate() {
	for _, t := range i {
		var err ErrPosition
		switch t.Type {

		// Let-operation
		case TOK_LET:
			//err = validate_let(i[1:])
		case TOK_COUT:
			err = validate_cout(i[1:], i[0])

		}

		if err.Error != nil {
			
			fmt.Printf("\t ... %s\n", t.Text)
			fmt.Printf("\t")
			for range err.Position + 5 {
				fmt.Print(" ")
			}
			fmt.Println("^")
			fmt.Printf("error: %s at position %d:%d\n", err.Error, err.Position, err.Line)
		}
	} 
}

func (p Program) Run() {
	for _, inst := range p.Instructions {

		inst.Validate()
	}
}

type Instruction []Token

type Program struct {
	Instructions []Instruction
}

func main() {
	fileName := flag.String("file", "", "path of the file to compile")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("please provide a filename using -file")
		os.Exit(1)
	}

	data, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("error opening file:", err)
		os.Exit(1)
	}

	sourceFile := SourceFile{
		FileName: *fileName,
		Content:  string(data),
	}

	var program Program

	lines := sourceFile.Lines()
	for _, l := range lines {
		tokens := l.Tokens()
		program.Instructions = append(program.Instructions, tokens)
	}

	program.Run()

}

/*

TODO:
-> Show Line, on which the error occured
-> show multiple arrows or the range of invalid arguments

*/
