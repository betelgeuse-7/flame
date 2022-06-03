package parser

import (
	"flame/scanner"
	"fmt"
	"testing"
)

func TestParserParseVarDeclStmt(t *testing.T) {
	input := "string name = \"Jennifer\""
	s := scanner.New(input)
	p := New(s)
	stmt := p.parseStmt()
	fmt.Println(stmt.String())
}
