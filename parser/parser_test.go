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
	input := `
		#string x = 56
		u32 y = true
		bool z = "Hey"
	`
	s := scanner.New(input)
	p := New(s)
	_ = p.ParseProgram()
	if errs := p.Errors(); len(errs) != 3 {
		t.Errorf("expected 3 errors in input code, but got %d errors\n", len(errs))
	}
	t.Logf("SUCCESS! Got 3 errors\n")
	for i, e := range p.Errors() {
		t.Logf("Err#%d: %s\n", i, e)
	}
}
