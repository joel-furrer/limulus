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
				Text:   scanner.Text(),
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
	Text   string
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
func validate_expression(i Instruction) ErrPosition {
	var errPos ErrPosition
	errPos.Line = i[0].Line
	var lastTok Token

	for j, t := range i {
		switch t.Type {
		case TOK_NUMBER:
			if lastTok.Type == TOK_NUMBER || lastTok.Type == TOK_IDENTIFIER {
				errPos.Error = fmt.Errorf("missing operand")
				errPos.Position = t.Position
				return errPos
			}

		case TOK_IDENTIFIER:
			if lastTok.Type == TOK_NUMBER || lastTok.Type == TOK_IDENTIFIER {
				errPos.Error = fmt.Errorf("missing operand")
				errPos.Position = t.Position
				return errPos
			}

			_, exists := getVar(t.Text)
			if !exists {
				errPos.Error = fmt.Errorf("undefined: %s", t.Text)
				errPos.Position = t.Position
			}

		case TOK_PLUS:
			if j == 0 {
				errPos.Error = fmt.Errorf("no value before '+'")
				errPos.Position = t.Position
				return errPos
			}

			if j == len(i)-1 {
				errPos.Error = fmt.Errorf("no value after '+'")
				errPos.Position = t.Position + 2
				return errPos
			}

			if lastTok.Type != TOK_NUMBER && lastTok.Type != TOK_IDENTIFIER {
				errPos.Error = fmt.Errorf("cannot use '+' for '%s'", lastTok.Text)
				errPos.Position = lastTok.Position
				return errPos
			}

		default:
			errPos.Error = fmt.Errorf("invalid token '%s' in expression", t.Text)
			errPos.Position = t.Position
			return errPos
		}
		lastTok = t
	}

	return errPos
}

func validate_let(i Instruction, t Token) ErrPosition {
	var errPos ErrPosition
	errPos.Line = t.Line

	// check for valid amount of arguments
	if len(i) < 3 {
		errPos.Error = fmt.Errorf("not enough arguments to call 'let'")
		errPos.Position = t.Position + len(t.Text) + 1
		return errPos
	}

	if i[1].Type != TOK_ASSIGN {
		errPos.Error = fmt.Errorf("missing '=' after let statement")
		errPos.Position = i[1].Position
		return errPos
	}

	setVar(i[0].Text, Variable{Type: "int"})

	// check expression
	expr := i[2:]
	errPos = validate_expression(expr)

	return errPos
}

func validate_cout(i Instruction, t Token) ErrPosition {
	var errPos ErrPosition
	errPos.Line = t.Line

	// check for valid amount of arguments
	if len(i) == 0 {
		errPos.Error = fmt.Errorf("missing argument for 'cout'")
		errPos.Position = t.Position + len(t.Text) + 1
		return errPos
	}

	// check for argument type
	expr := i
	errPos = validate_expression(expr)

	return errPos
}

type ErrPosition struct {
	Error    error
	Line     int
	Position int
}

func (p Program) Validate() bool {
	for _, i := range p.Instructions {
		t := i[0]
		var err ErrPosition
		switch t.Type {

		// Let-operation
		case TOK_LET:
			err = validate_let(i[1:], i[0])
		case TOK_COUT:
			err = validate_cout(i[1:], i[0])

		}

		if err.Error != nil {

			i.Print()
			for range err.Position {
				fmt.Print(" ")
			}
			fmt.Println("^")
			fmt.Printf("./%s:%d:%d: %s\n", p.Name, err.Position, err.Line, err.Error)
			return false
		}
	}
	return true
}

func (i Instruction) Print() {
	posProg := 0
	tokIndex := 0
	lastTok := i[len(i)-1]
	lastPos := lastTok.Position + len(lastTok.Text)

	for range lastPos {
		if posProg == i[tokIndex].Position {
			fmt.Print(i[tokIndex].Text)
			posProg += len(i[tokIndex].Text)
			if tokIndex < len(i)-1 {
				tokIndex++
			}
		} else {
			fmt.Print(" ")
			posProg++
		}
	}

	fmt.Println()
}

type Instruction []Token

type Program struct {
	Instructions []Instruction
	Name         string
}

// variables
type Variable struct {
	Type string
	Data interface{}
}

func setVar(name string, v Variable) {
	variables[name] = v
}

func getVar(name string) (Variable, bool) {
	v, exists := variables[name]
	return v, exists
}

var variables = make(map[string]Variable)

func (p Program) ParseToAST() {
	for _, i := range p.Instructions {
		t := i[0]
		switch t.Type {
		case TOK_LET:
			var assignNode AssignmentNode
			assignNode.Name = i[1].Text
			if len(i) == 4 {
				switch i[3].Type {
				case TOK_NUMBER:
					v, _ := strconv.Atoi(i[3].Text)
					numNode := NumberNode{Value: v }
					assignNode.Value = numNode
				case TOK_IDENTIFIER:
					idNode := IdentifierNode{Name: i[3].Text }
					assignNode.Value = idNode
				}
			} else {
				exprAst, err := i[3:].ParseExpression()
				if err != nil {
					fmt.Printf("error parsing expression: %v\n", err)
				}
				assignNode.Value = exprAst
			}

			assignNode.Print()
		}
	}
}

func (a AssignmentNode) Print() {
    fmt.Println("Assignment")
    fmt.Printf("    Identifier: %s\n", a.Name)
	printNode(a.Value, "    ")
	fmt.Println()
}

func printNode(n Node, indent string) {
    switch v := n.(type) {
    case NumberNode:
        fmt.Printf("%sLiteral: %d\n", indent, v.Value)
    case IdentifierNode:
        fmt.Printf("%sIdentifier: %s\n", indent, v.Name)
    case BinOpNode:
        fmt.Printf("%sBinaryOp (%s)\n", indent, v.Operator)
        printNode(v.Left, indent+"    ")
        printNode(v.Right, indent+"    ")
    default:
        fmt.Printf("%sUnknown node\n", indent)
    }
}

func StrToInt(s string) (int, error) {
    value, err := strconv.Atoi(s)
    if err != nil {
        return 0, fmt.Errorf("cannot convert %q to int: %w", s, err)
    }
    return value, nil
}

func (i Instruction) ParseExpression() (Node, error) {
    if len(i) < 3 {
        return nil, fmt.Errorf("expression too short")
    }

    var left Node
    first := i[0]
    switch first.Type {
    case TOK_NUMBER:
        v, _ := strconv.Atoi(first.Text)
        left = NumberNode{Value: v}
    case TOK_IDENTIFIER:
        left = IdentifierNode{Name: first.Text}
    default:
        return nil, fmt.Errorf("unexpected token %v", first)
    }

    pos := 1
    for pos < len(i) {
        opTok := i[pos]
        if opTok.Type != TOK_PLUS {
            break
        }
        pos++

        if pos >= len(i) {
            return nil, fmt.Errorf("expected right operand after operator")
        }
        rightTok := i[pos]
        var right Node
        switch rightTok.Type {
        case TOK_NUMBER:
            v, _ := strconv.Atoi(rightTok.Text)
            right = NumberNode{Value: v}
        case TOK_IDENTIFIER:
            right = IdentifierNode{Name: rightTok.Text}
        default:
            return nil, fmt.Errorf("unexpected token %v", rightTok)
        }

        left = BinOpNode{
            Left:     left,
            Operator: opTok.Text,
            Right:    right,
        }

        pos++
    }

    return left, nil
}

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

type NumberNode struct {
	Value int
}

type IdentifierNode struct {
	Name string
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
	program.Name = sourceFile.FileName

	lines := sourceFile.Lines()
	for _, l := range lines {
		tokens := l.Tokens()
		program.Instructions = append(program.Instructions, tokens)
	}

	valid := program.Validate()
	if !valid {
		os.Exit(1)
	}

	// create AST from valid program
	program.ParseToAST()
}
