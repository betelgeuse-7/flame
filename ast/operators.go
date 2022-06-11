package ast

import "flame/token"

type Operator uint32

const (
	OP_ADD Operator = iota + 1
	OP_SUB
	OP_MUL
	OP_DIV
	OP_MOD
)

// return token struct of an operator
func (o Operator) Token() (tok token.Token) {
	tok = token.Token{}
	switch o {
	case OP_ADD:
		tok.Typ = token.T_Plus
		tok.Lit = "+"
	case OP_SUB:
		tok.Typ = token.T_Minus
		tok.Lit = "-"
	case OP_MUL:
		tok.Typ = token.T_Mul
		tok.Lit = "*"
	case OP_DIV:
		tok.Typ = token.T_Div
		tok.Lit = "/"
	case OP_MOD:
		tok.Typ = token.T_Modulus
		tok.Lit = "%"
	}
	return tok
}
