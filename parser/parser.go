package parser

import (
	"flame/ast"
	"flame/scanner"
	"flame/token"
	"strconv"
)

// TODO Fix line numbers in parser errors.

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
	for p.cur.Typ != token.Eof {
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
	if p.cur.Typ == token.Eof {
		return nil
	}
	switch p.cur.Typ {
	case token.Octothorp:
		return p.parseConstDecl()
	case token.If:
		return p.parseIfStmt()
	case token.Else, token.ElseIf:
		p.reportErr("no 'elseif's, or 'else's without a preceding if")
		return nil
	default:
		return p.parseExprStmt()
	}
}

// TODO var && const decl parsing
// types can also be slices, maps, ...
// value can be another variable/const, or a function call, or a slice/map literal, etc.

func (p *Parser) parseVarDecl() *ast.VarDeclStmt {
	s := &ast.VarDeclStmt{DataType: p.cur.Typ}
	if ok := p.assertPeek(token.Ident); !ok {
		return nil
	}
	s.Name = p.cur.Lit
	if ok := p.assertPeek(token.Eq); !ok {
		return nil
	}
	p.advance()
	s.Value = p.parseExpr()
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
	if ok := p.assertPeek(token.Ident); !ok {
		return nil
	}
	s.Decl.Name = p.cur.Lit
	if ok := p.assertPeek(token.Eq); !ok {
		return nil
	}
	p.advance()
	s.Decl.Value = p.parseExpr()
	return s
}

func (p *Parser) parseExprStmt() *ast.ExprStmt {
	stmt := &ast.ExprStmt{}
	stmt.Expr = p.parseExpr()
	return stmt
}

func (p *Parser) parseExpr() ast.Expr {
	// TODO error handling
	var leftExpr ast.Expr
	switch p.cur.Typ {
	case token.Uint:
		leftExpr = p.parseUnsignedIntegerLiteral()
	case token.Int:
		leftExpr = p.parseSignedIntegerLiteral()
	case token.Float:
		leftExpr = p.parseFloatLiteral()
	case token.Bool:
		leftExpr = p.parseBoolLiteral()
	case token.String:
		leftExpr = p.parseStringLiteral()
	}
	// is this an infix expr, or a prefix expr ?
	infixFn := p.infixParseFns[p.peek.Typ]
	if infixFn == nil {
		return leftExpr
	}
	p.advance()
	expr := infixFn(leftExpr)
	return expr
}

func (p *Parser) parseStringLiteral() *ast.StringLiteral {
	return &ast.StringLiteral{Val: p.cur.Lit}
}

func (p *Parser) parseSignedIntegerLiteral() ast.SignedIntegerLiteral {
	val, err := strconv.ParseInt(p.cur.Lit, 10, 64)
	if err != nil {
		return nil
	}
	return &ast.IntLiteral{ValStr: p.cur.Lit, Val: val}
}

func (p *Parser) parseUnsignedIntegerLiteral() ast.UnsignedIntegerLiteral {
	val, err := strconv.ParseUint(p.cur.Lit, 10, 64)
	if err != nil {
		return nil
	}
	return &ast.UintLiteral{ValStr: p.cur.Lit, Val: val}
}

func (p *Parser) parseFloatLiteral() ast.IFloatLiteral {
	val, err := strconv.ParseFloat(p.cur.Lit, 64)
	if err != nil {
		return nil
	}
	return &ast.FloatLiteral{ValStr: p.cur.Lit, Val: val}
}

func (p *Parser) parseBoolLiteral() *ast.BooleanLiteral {
	val, err := strconv.ParseBool(p.cur.Lit)
	if err != nil {
		return nil
	}
	return &ast.BooleanLiteral{ValStr: p.cur.Lit, Val: val}
}

func (p *Parser) parseInfixExpr(left ast.Expr) ast.Expr {
	expr := &ast.BinOp{
		Lhs:      left,
		Operator: p.cur.Lit,
	}
	p.advance()
	expr.Rhs = p.parseExpr()
	if expr.Rhs == nil {
		x := p.cur.Pos.X
		y := p.cur.Pos.Y
		p.reportErr("%d:%d Expected another expression after '%s', but got nothing", y, x, expr.Operator)
		return nil
	}
	return expr
}

func (p *Parser) parseIfStmt() *ast.IfStmt {
	stmt := &ast.IfStmt{}
	p.advance()
	stmt.Cond = p.parseExpr()
	if stmt.Cond == nil {
		p.reportErr("empty condition in if statement")
		return nil
	}
	if ok := p.assertPeek(token.LCurly); !ok {
		return nil
	}
	for p.peek.Typ != token.RCurly {
		p.advance()
		if p.cur.Typ == token.Eof {
			p.reportErr("unexpected EOF in if statement body")
			return nil
		}
		if newStmt := p.parseStmt(); newStmt != nil {
			stmt.Body = append(stmt.Body, newStmt)
		}
	}
	if ok := p.assertPeek(token.RCurly); !ok {
		return nil
	}
	switch p.peek.Typ {
	case token.ElseIf:
		p.advance()
		stmt.Alternative = p.parseIfStmt()
	case token.Else:
		p.advance()
		elseBranchStmts := []ast.Stmt{}
		if ok := p.assertPeek(token.LCurly); !ok {
			return nil
		}
		for p.peek.Typ != token.RCurly {
			p.advance()
			if p.cur.Typ == token.Eof {
				p.reportErr("unexpected EOF in else statement body")
				return nil
			}
			if newElseBranchStmt := p.parseStmt(); newElseBranchStmt != nil {
				elseBranchStmts = append(elseBranchStmts, newElseBranchStmt)
			}
		}
		if ok := p.assertPeek(token.RCurly); !ok {
			return nil
		}
		stmt.Default = elseBranchStmts
		p.advance()
	}
	return stmt
}
