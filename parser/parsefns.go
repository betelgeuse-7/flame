package parser

import (
	"flame/ast"
	"flame/token"
)

type infixParseFn func(ast.Expr) ast.Expr

func (p *Parser) registerAllExprParseFns() {
	/* INFIX */
	//p.registerInfixFn(token.T_Ident, p.parseInfixExpr)
	p.registerInfixFn(token.T_Plus, p.parseInfixExpr)
	p.registerInfixFn(token.T_Minus, p.parseInfixExpr)
	p.registerInfixFn(token.T_Div, p.parseInfixExpr)
	p.registerInfixFn(token.T_Mul, p.parseInfixExpr)
	p.registerInfixFn(token.T_Modulus, p.parseInfixExpr)
	p.registerInfixFn(token.T_Lt, p.parseInfixExpr)
	p.registerInfixFn(token.T_LtEq, p.parseInfixExpr)
	p.registerInfixFn(token.T_Gt, p.parseInfixExpr)
	p.registerInfixFn(token.T_GtEq, p.parseInfixExpr)
	p.registerInfixFn(token.T_NotEq, p.parseInfixExpr)
	p.registerInfixFn(token.T_DoubleEq, p.parseInfixExpr)
	p.registerInfixFn(token.T_And, p.parseInfixExpr)
	p.registerInfixFn(token.T_Or, p.parseInfixExpr)
	p.registerInfixFn(token.T_BitLeftShift, p.parseInfixExpr)
	p.registerInfixFn(token.T_BitRightShift, p.parseInfixExpr)
	//p.registerInfixFn(token.T_SingleQuote, p.parseInfixExpr)
}

func (p *Parser) registerInfixFn(tt token.TokenType, fn infixParseFn) {
	p.infixParseFns[tt] = fn
}
