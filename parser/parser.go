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
		return p.parseVarDeclStmt(false, false)
	case token.Octothorp:
		// if next token is a [, or a {; then it must be a compound type
		next := p.peek.Typ
		if next == token.LSquareParen || next == token.LCurly {
			// starting token = #
			return p.parseCompoundConstDecl()
		}
		return p.parseConstDeclStmt(false, false)
	case token.LSquareParen, token.LCurly:
		return p.parseCompoundVarDecl()
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

// TODO: parse map, and slice literals

func (p *Parser) parseConstDeclStmt(isSlice, isMap bool) ast.Stmt {
	if isSlice && isMap {
		panic("parseConstDeclStmt: isMap, and isSlice are both true")
	}
	tok := p.cur.Typ
	tokPos := p.cur.Pos
	p.advance()
	// expect a type after #
	var dataType ast.Type
	if isSlice {
		p.advance()
		dataType = p.decideDataType(p.cur.Lit)
		dataType = ast.SliceType{
			Typ: dataType,
		}
		if ok := p.expect(token.RSquareParen); !(ok) {
			return nil
		}
	} else if isMap {
		p.advance()
		keyType := p.decideDataType(p.cur.Lit)
		if ok := p.expect(token.Colon); !(ok) {
			return nil
		}
		peekTyp := p.peek.Typ
		if peekTyp == token.Eof || peekTyp == token.Illegal || peekTyp == token.Newline {
			p.reportErr("expected another type for value in map type declaration")
			return nil
		}
		p.advance()
		valType := p.decideDataType(p.cur.Lit)
		dataType = ast.MapType{
			Key:   keyType,
			Value: valType,
		}
		if ok := p.expect(token.RCurly); !(ok) {
			return nil
		}
	} else {
		if ok := p.expectType(true); !(ok) {
			return nil
		}
		dataType = p.decideDataType(p.cur.Lit)
	}
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
		p.reportErr("missing expression in constant declaration")
		return nil
	}
	stmt := &ast.ConstDeclStmt{}
	stmt.Type = dataType
	stmt.Ident = ident.(ast.Ident)
	stmt.Tok = tok
	stmt.TokPos = tokPos
	stmt.Value = value
	return stmt
}

func (p *Parser) parseVarDeclStmt(isSlice, isMap bool) ast.Stmt {
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

func (p *Parser) parseCompoundConstDecl() ast.Stmt {
	isSlice, isMap := p.peek.Typ == token.LSquareParen, p.peek.Typ == token.LCurly
	return p.parseConstDeclStmt(isSlice, isMap)
}

func (p *Parser) parseCompoundVarDecl() ast.Stmt {
	isSlice, isMap := p.peek.Typ == token.LSquareParen, p.peek.Typ == token.LCurly
	return p.parseVarDeclStmt(isSlice, isMap)
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
		token.Int, token.Float, token.Bool, token.String, token.Char,
	}
	return in(toks, cur)
}

func (p *Parser) decideDataType(lit string) ast.Type {
	/*
		? when extending this method, we will need to support user-defined types.
		? those will probably be called ast.StructType.
		=============================
		> for user-defined types, the key will probably be "IDENT", or "ident".
		> I am not sure.
	*/
	m := map[string]ast.Type{
		"string": ast.StringType{},
		"uint":   ast.UintType{},
		"int":    ast.IntType{},
		"float":  ast.FloatType{},
		"bool":   ast.BoolType{},
		"char":   ast.CharType{},
	}
	val := m[lit]
	// nil if non-existent
	return val
}
