package main

import (
	"flag"
	"fmt"
	"os"

	"limulus/lexer"
	"limulus/parser"
)

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
	}

	source := lexer.SourceFile{
		Name:    *fileName,
		Content: string(data),
	}

	// lexer
	instructions := lexer.Lex(source)

	// parser
	_, err = parser.Parse(instructions)
	if err != nil {
		fmt.Println("parse error:", err)
		os.Exit(1)
	}

	//fmt.Println(program)
}
