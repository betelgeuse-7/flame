package parser

import (
	"flame/ast"
	"flame/scanner"
	"flame/token"
	"fmt"
	"reflect"
)

type Parser struct {
	scanner *scanner.Scanner
	cur     token.Token
	peek    token.Token
	errors  []string

	// infix expression parsing methods
	infixParseFns map[token.TokenType]infixParseFn
}

func New(scanner *scanner.Scanner) *Parser {
	p := &Parser{
		scanner:       scanner,
		infixParseFns: make(map[token.TokenType]infixParseFn),
	}
	p.registerAllExprParseFns()
	p.advance()
	p.advance()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.peek.Typ != token.Eof {
		stmt := p.parseStmt()
		fmt.Println("ParseProgram: reflect: ", reflect.TypeOf(stmt))
		if stmt != nil {
			program.Stmts = append(program.Stmts, stmt)
		}
		p.advance()
	}
	return program
}

func (p *Parser) parseStmt() ast.Stmt {
	if p.cur.Typ == token.Eof {
		return nil
	}
	for p.cur.Typ == token.Newline {
		p.advance()
	}
	switch p.cur.Typ {
	case token.BoolKw, token.FloatKw, token.StringKw, token.IntKw, token.UintKw, token.Octothorp:
		return p.parseGenDeclStmt()
	}
	if p.shouldCallParseExprStmt() {
		return p.parseExprStmt()
	}
	p.reportErr("unexpected token: %s", p.cur.Lit)
	return nil
}

func (p *Parser) parseGenDeclStmt() ast.Stmt {
	isConst := p.cur.Typ == token.Octothorp
	if isConst {
		p.advance()
	}
	tok := p.cur.Typ
	tokPos := p.cur.Pos
	// parse identifier
	if ok := p.expect(token.Ident); !(ok) {
		return nil
	}
	ident := p.parseIdent()
	if ok := p.expect(token.Eq); !(ok) {
		return nil
	}
	if p.peek.Typ == token.Eof {
		p.reportErr("unexpected EOF")
		return nil
	}
	p.advance()
	value := p.parseExpr()
	genDeclStmt := ast.GenericDeclStmt{
		TokPos: tokPos,
		Tok:    tok,
		Ident:  ident.(ast.Ident),
		Value:  value,
	}
	if isConst {
		return ast.ConstDeclStmt{
			GenericDeclStmt: genDeclStmt,
		}
	}
	return ast.VarDeclStmt{
		GenericDeclStmt: genDeclStmt,
	}
}

func (p *Parser) shouldCallParseExprStmt() bool {
	cur := p.cur.Typ
	in := func(tt []token.TokenType, t token.TokenType) bool {
		for _, v := range tt {
			if t == v {
				return true
			}
		}
		return false
	}
	toks := []token.TokenType{
		token.Plus, token.Minus, token.Div, token.Mul, token.Lt, token.LtEq, token.Modulus,
		token.Gt, token.GtEq, token.NotEq, token.DoubleEq, token.And, token.Or, token.Uint,
		token.Int, token.Float, token.Bool, token.String,
	}
	return in(toks, cur)
}
