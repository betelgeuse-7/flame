package parser

import (
	"flame/ast"
	"flame/token"
	"fmt"
	"strconv"
)

func (p *Parser) parseExprStmt() *ast.ExprStmt {
	stmt := &ast.ExprStmt{}
	stmt.Expr = p.parseExpr()
	return stmt
}

func (p *Parser) parseExpr() ast.Expr {
	fmt.Println("parseExpr: ", p.cur.Typ)
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
	case token.Ident:
		// parse ident
		leftExpr = p.parseIdent()
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

func (p *Parser) parseIdent() ast.Expr {
	return ast.Ident{Pos: p.cur.Pos, Ident: p.cur.Lit}
}
