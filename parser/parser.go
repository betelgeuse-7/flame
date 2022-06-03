package parser

import (
	"flame/ast"
	"flame/scanner"
	"flame/token"
	"fmt"
)

type Parser struct {
	scanner *scanner.Scanner
	cur     token.Token
	peek    token.Token
	errors  []string
}

func New(scanner *scanner.Scanner) *Parser {
	p := &Parser{
		scanner: scanner,
	}
	p.advance()
	p.advance()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

// ! ignore token.T_Whitespace
// ? should we just ignore whitespaces in lexical analysis? (in scanner)

func (p *Parser) advance() {
	p.cur = p.peek
	p.peek = p.scanner.Next()
	for p.cur.Typ == token.T_Whitespace {
		p.cur = p.peek
		p.peek = p.scanner.Next()
	}
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peek.Typ == t {
		p.advance()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peek.Typ)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Stmts = []ast.Stmt{}
	for p.cur.Typ != token.T_Eof {
		stmt := p.parseStmt()
		if stmt != nil {
			program.Stmts = append(program.Stmts, stmt)
		}
		p.advance()
	}
	return program
}

func (p *Parser) parseStmt() ast.Stmt {
	return p.parseVarDecl()
	/*
		if isDataTypeKw(p.cur.Typ) {
			return p.parseVarDecl()
		}
		panic("*Parser.parseStmt: not implemented")
	*/
}

func (p *Parser) parseVarDecl() ast.Stmt {
	stmt := &ast.VariableDeclarationStmt{DataType: p.cur}
	if !p.expectPeek(token.T_Ident) {
		fmt.Printf("*Parser.parseVarDecl: p.expectPeek#1: p.peek = %s\n", p.peek)
		return nil
	}
	stmt.Name = p.cur.Lit
	if !p.expectPeek(token.T_Eq) {
		fmt.Printf("*Parser.parseVarDecl: p.expectPeek#2: p.peek = %s\n", p.peek)
		return nil
	}
	p.advance()
	stmt.Value = p.cur.Lit
	return stmt
}

func isDataTypeKw(typ token.TokenType) bool {
	for _, v := range dataTypeKws {
		if typ == v {
			return true
		}
	}
	return false
}
