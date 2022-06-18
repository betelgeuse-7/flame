package parser

import (
	"flame/scanner"
	"strings"
	"testing"
)

func TestParserParseProgram(t *testing.T) {
	input := `#uint age = 44
		      #bool isRaining = true
			  #f32 PI = 3.14
	          #string name = "Jennifer"
			  string city = "Amsterdam"
			  float temp = 29.8`
	s := scanner.New(input)
	p := New(s)
	program := p.ParseProgram()
	if errs := p.errors; len(errs) > 0 {
		for _, v := range errs {
			t.Logf("Parser error: %s\n", v)
		}
		t.Fatalf("there were parser errors\n")
	}
	lines := strings.Split(input, "\n")
	lenLines := len(lines)
	if len(program.Stmts) != lenLines {
		t.Errorf("expected %d statements, got %d statements", lenLines, len(program.Stmts))
		t.Logf("Statements: %s", program.Stmts)
	}
	t.Logf("SUCCESS! Statements: %s", program.Stmts)
}

func TestParserParseProgramErrors(t *testing.T) {
	expectedErrCount := 6
	input := "#string x = 56\n"
	input += "u32 y = true\n"
	input += "bool z = \"Hey\"\n"
	input += "#f32 s 8.16\n"
	input += "#string pp =\n"
	input += "#u32 oo =\n"

	s := scanner.New(input)
	p := New(s)
	_ = p.ParseProgram()
	if errs := p.Errors(); len(errs) != expectedErrCount {
		t.Fatalf("expected %d errors in input code, but got %d errors\n", expectedErrCount, len(errs))
	}
	t.Logf("SUCCESS! Got %d errors\n", expectedErrCount)
	for i, e := range p.Errors() {
		t.Logf("Err#%d: %s\n", i, e)
	}
}

func TestParserParseBinOpExprStmt(t *testing.T) {
	input := "#int x = 5 + 3\n"
	input += "u32 y = 55 + -16 * 558\n"
	input += "5 +\n"
	input += "4 != 6 \"a\" && \"b\" \n"
	input += "5 -\n"
	input += "2 << 1\n"
	input += "-16.7 >> \n"

	s := scanner.New(input)
	p := New(s)
	program := p.ParseProgram()
	t.Logf("STATEMENTS: %s", program.String())
	t.Logf("errors len: %d\n", len(p.errors))
	t.Logf("Parser errors: %v", p.Errors())
}

func TestParserParseIfStmt(t *testing.T) {
	input := "if true { #string x = \"x\"} elseif false { } else { }\n"
	s := scanner.New(input)
	p := New(s)
	program := p.ParseProgram()
	if len(p.errors) > 0 {
		t.Logf("errors: %v\n", p.errors)
		return
	}
	t.Logf("STATEMENTS: %v\n", program.Stmts)
}
