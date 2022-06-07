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

func (p *Parser) advance() {
	p.cur = p.peek
	p.peek = p.scanner.Next()
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

func (p *Parser) reportErr(format string, args ...interface{}) {
	p.errors = append(p.errors, fmt.Sprintf(format, args...))
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
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
	if isDataTypeKw(p.cur.Typ) {
		return p.parseVarDecl()
	}
	switch p.cur.Typ {
	case token.T_Octothorp:
		return p.parseConstDecl()
	}
	return nil
}

func (p *Parser) parseVarDecl() *ast.VarDeclStmt {
	s := &ast.VarDeclStmt{DataType: p.cur.Typ}
	if ok := p.expectPeek(token.T_Ident); !ok {
		return nil
	}
	s.Name = p.cur.Lit
	if ok := p.expectPeek(token.T_Eq); !ok {
		return nil
	}
	p.advance()
	if ok := checkPrimitiveValue(p, s.DataType); !ok {
		return nil
	}
	s.Value = ast.PrimitiveValue{DataType: s.DataType, Val: p.cur.Lit}
	return s
}

func (p *Parser) parseConstDecl() *ast.ConstDeclStmt {
	s := &ast.ConstDeclStmt{Octothorp: p.cur.Typ}
	if !(isDataTypeKw(p.peek.Typ)) {
		p.reportErr("Malformed Constant Declaration: expected a data type keyword after an octothorp, got %s", p.peek.Lit)
		return nil
	}
	p.advance()
	s.Decl.DataType = p.cur.Typ
	if ok := p.expectPeek(token.T_Ident); !ok {
		return nil
	}
	s.Decl.Name = p.cur.Lit
	if ok := p.expectPeek(token.T_Eq); !ok {
		return nil
	}
	p.advance()
	if ok := checkPrimitiveValue(p, s.Decl.DataType); !ok {
		return nil
	}
	s.Decl.Value = ast.PrimitiveValue{DataType: s.Decl.DataType, Val: p.cur.Lit}
	return s
}

func isDataTypeKw(typ token.TokenType) bool {
	for _, v := range dataTypeKws {
		if typ == v {
			return true
		}
	}
	return false
}
