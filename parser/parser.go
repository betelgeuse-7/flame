package parser

import (
	"flame/ast"
	"flame/scanner"
	"flame/token"
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
		if stmt != nil {
			program.Stmts = append(program.Stmts, stmt)
		}
		p.advance()
	}
	return program
}

func (p *Parser) parseStmt() ast.Stmt {
	for p.cur.Typ == token.Newline {
		p.advance()
	}
	switch p.cur.Typ {
	case token.BoolKw, token.FloatKw, token.StringKw, token.IntKw, token.UintKw:
		return p.parseVarDeclStmt()
	case token.Octothorp:
		return p.parseConstDeclStmt()
	}
	if p.shouldCallParseExprStmt() {
		return p.parseExprStmt()
	}
	if p.cur.Typ == token.Eof {
		return nil
	}
	p.reportErr("unexpected token: %s", p.cur.Lit)
	return nil
}

func (p *Parser) parseConstDeclStmt() ast.Stmt {
	// expect a type after #
	illegal, newline, eof := token.Illegal, token.Newline, token.Eof
	p.advance()
	curTyp := p.cur.Typ
	if curTyp == illegal || curTyp == newline || curTyp == eof {
		p.reportErr("expected a type after '#'")
		return nil
	}
	tok := p.cur.Typ
	tokPos := p.cur.Pos
	if ok := p.expect(token.Ident); !(ok) {
		return nil
	}
	ident := p.parseIdent()
	if ok := p.expect(token.Eq); !(ok) {
		return nil
	}
	p.advance()
	if p.cur.Typ == eof {
		p.reportErr("unexpected EOF")
		return nil
	}
	value := p.parseExpr()
	if value == nil {
		p.reportErr("missing expression in constant declaration")
		return nil
	}
	stmt := &ast.ConstDeclStmt{}
	stmt.Ident = ident.(ast.Ident)
	stmt.Tok = tok
	stmt.TokPos = tokPos
	stmt.Value = value
	return stmt
}

func (p *Parser) parseVarDeclStmt() ast.Stmt {
	tok := p.cur.Typ
	tokPos := p.cur.Pos
	if ok := p.expect(token.Ident); !(ok) {
		return nil
	}
	ident := p.parseIdent()
	if ok := p.expect(token.Eq); !(ok) {
		return nil
	}
	p.advance()
	if p.cur.Typ == token.Eof {
		p.reportErr("unexpected EOF")
		return nil
	}
	value := p.parseExpr()
	if value == nil {
		p.reportErr("missing expression in variable declaration")
		return nil
	}
	stmt := &ast.VarDeclStmt{}
	stmt.Ident = ident.(ast.Ident)
	stmt.Tok = tok
	stmt.TokPos = tokPos
	stmt.Value = value
	return stmt
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
