package parser

import (
	"flame/ast"
	"flame/token"
)

type infixParseFn func(ast.Expr) ast.Expr

func (p *Parser) registerAllExprParseFns() {
	/* INFIX */
	//p.registerInfixFn(token.Ident, p.parseInfixExpr)
	p.registerInfixFn(token.Plus, p.parseInfixExpr)
	p.registerInfixFn(token.Minus, p.parseInfixExpr)
	p.registerInfixFn(token.Div, p.parseInfixExpr)
	p.registerInfixFn(token.Mul, p.parseInfixExpr)
	p.registerInfixFn(token.Modulus, p.parseInfixExpr)
	p.registerInfixFn(token.Lt, p.parseInfixExpr)
	p.registerInfixFn(token.LtEq, p.parseInfixExpr)
	p.registerInfixFn(token.Gt, p.parseInfixExpr)
	p.registerInfixFn(token.GtEq, p.parseInfixExpr)
	p.registerInfixFn(token.NotEq, p.parseInfixExpr)
	p.registerInfixFn(token.DoubleEq, p.parseInfixExpr)
	p.registerInfixFn(token.And, p.parseInfixExpr)
	p.registerInfixFn(token.Or, p.parseInfixExpr)
	//p.registerInfixFn(token.SingleQuote, p.parseInfixExpr)
}

func (p *Parser) registerInfixFn(tt token.TokenType, fn infixParseFn) {
	p.infixParseFns[tt] = fn
}
