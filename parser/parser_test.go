package parser

import (
	"flame/scanner"
	"fmt"
	"testing"
)

func TestParserParseGenDeclStmt(t *testing.T) {
	input := `
		string x = "hey"
		#string y = "hello"
		int n = -16
		uint nn = 1576
		#float PI = 3.14
		bool isRaining = false
		int n2 = n
		.
	`
	s := scanner.New(input)
	p := New(s)
	parsed := p.ParseProgram()
	fmt.Println(len(parsed.Stmts))
	fmt.Println(parsed.Stmts)
	errs := p.errors
	if len(errs) > 0 {
		fmt.Println("errors: ", errs)
		return
	}
	/*
		for _, v := range parsed.Stmts {
			fmt.Printf(">>> %s\n", v.String())
		}*/
}

func TestParserParseBinOpExprStmt(t *testing.T) {
	input := "#int x = 5 + 3\n"
	input += "u32 y = 55 + -16 * 558\n"
	input += "5 +\n"
	input += "4 != 6 \"a\" && \"b\" \n"
	input += "5 -\n"
	input += "2 + 1\n"
	input += "-16.7 - 1\n"
	input += "#f32 n = foo"

	s := scanner.New(input)
	p := New(s)
	program := p.ParseProgram()
	for _, v := range program.Stmts {
		fmt.Println("stmt >>>>", v)
	}
	/*
		t.Logf("STATEMENTS: %s", program.Stmts)
		t.Logf("errors len: %d\n", len(p.errors))
		t.Logf("Parser errors: %v", p.Errors())
	*/
}
