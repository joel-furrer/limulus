package lexer

import (
	"bufio"
	"strings"
)

type SourceFile struct {
	Name    string
	Content string
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
