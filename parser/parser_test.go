package parser

import (
	"flame/ast"
	"flame/scanner"
	"fmt"
	"testing"
)

func _errPrint(errs []string) {
	for _, v := range errs {
		fmt.Printf("[!!!] > %s\n", v)
	}
}

func _stmtPrint(stmts []ast.Stmt) {
	for _, v := range stmts {
		fmt.Printf(">>> %+v\n", v)
	}
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
	_errPrint(p.errors)
	_stmtPrint(program.Stmts)
}

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
		int
		int x 
		int x = 
		#string
	`
	s := scanner.New(input)
	p := New(s)
	parsed := p.ParseProgram()
	errs := p.errors
	_errPrint(errs)
	_stmtPrint(parsed.Stmts)
}

// test var/const decls with compound types
func TestParserParseGenDeclStmt2(t *testing.T) {
	input := `
		#[string] nx = ["hey", "hello", "ya", "j"]
	`
	s := scanner.New(input)
	p := New(s)
	parsed := p.ParseProgram()
	_errPrint(p.errors)
	_stmtPrint(parsed.Stmts)
}

func TestParserParseGenDeclStmt3(t *testing.T) {
	input := `
		#[string] nx = ["hey", ]
	`
	s := scanner.New(input)
	p := New(s)
	parsed := p.ParseProgram()
	_errPrint(p.errors)
	_stmtPrint(parsed.Stmts)
}

func TestParserParseGenDeclStmt4(t *testing.T) {
	input := `
		#[string] nx = ["hey" "hello"]
	`
	s := scanner.New(input)
	p := New(s)
	parsed := p.ParseProgram()
	_errPrint(p.errors)
	_stmtPrint(parsed.Stmts)
}

func TestParserParseGenDeclStmt5(t *testing.T) {
	input := `
		#[string] a = ["ye", "s", "!"]
		#[float] fx = [3.14, 5.17, 5.1]
	`
	s := scanner.New(input)
	p := New(s)
	parsed := p.ParseProgram()
	_errPrint(p.errors)
	_stmtPrint(parsed.Stmts)
}

func TestParserParseGenDeclStmt6(t *testing.T) {
	input := `
		#[string] a = ["a" "b" "c""d" "e" "f"]
		#[string] b = ["a","b","c","d","e","f"]
		#[string]c=["a","b","c",,,,,,,,,,,"d","e","f",,,,]
`
	s := scanner.New(input)
	p := New(s)
	parsed := p.ParseProgram()
	_errPrint(p.errors)
	_stmtPrint(parsed.Stmts)
}

func TestParserParseGenDeclStmt7(t *testing.T) {
	input := `
		#[string] a = ["a","b","c","d","e","f"]
`
	s := scanner.New(input)
	p := New(s)
	parsed := p.ParseProgram()
	_errPrint(p.errors)
	_stmtPrint(parsed.Stmts)
}

func TestParserParseMapLit(t *testing.T) {
	input := `
		#{string:string} px = {"name": "Jennifer", "age": "67"}
	`
	s := scanner.New(input)
	p := New(s)
	program := p.ParseProgram()
	_errPrint(p.errors)
	_stmtPrint(program.Stmts)
}
