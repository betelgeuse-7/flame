package parser

import (
	"flame/ast"
	"flame/scanner"
	"flame/token"
	"fmt"
)

// TODO Fix line numbers in parser errors.

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
	msg := fmt.Sprintf("%d:%d Expected next token to be %s, got %s instead", p.cur.Pos.Y, p.cur.Pos.X, t, p.peek.Typ)
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
	default:
		return p.parseExprStmt()
	}
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
	val, err := giveProperValue(s.DataType, p.cur.Lit)
	if err != nil {
		return nil
	}
	s.Value = val
	return s
}

func (p *Parser) parseConstDecl() *ast.ConstDeclStmt {
	s := &ast.ConstDeclStmt{Octothorp: p.cur.Typ}
	if !(isDataTypeKw(p.peek.Typ)) {
		p.reportErr("%d:%d Malformed Constant Declaration: expected a data type keyword after an octothorp, got %s", p.cur.Pos.Y, p.cur.Pos.X, p.peek.Lit)
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
	val, err := giveProperValue(s.Decl.DataType, p.cur.Lit)
	if err != nil {
		return nil
	}
	s.Decl.Value = val
	return s
}

func (p *Parser) parseExprStmt() *ast.ExprStmt {
	// PREFIX: 	++ + -- - ( * & ...
	// INFIX: 	+ - / * % ...
	// POSTFIX: ++ -- ' ...
	return nil
}
