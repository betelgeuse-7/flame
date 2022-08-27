package ast

import (
	"fmt"
)

type Program struct {
	Stmts []Stmt
}

/*
func (p *Program) String() string {
	res := ""
	for i, v := range p.Stmts {
		res += fmt.Sprintf("(%s)  ", v.String())
		if i%3 == 0 {
			res += "\n"
		}
	}
	return res
}
*/
type Node interface {
	String() string
}

type Stmt interface {
	Node
	S()
}

type Expr interface {
	Node
	E()
}

type ExprStmt struct {
	Expr
}

func (e ExprStmt) S() {}
func (e ExprStmt) String() string {
	return e.Expr.String()
}

type PrefixOp struct {
	Operator string
	Rhs      Expr
}

func (p PrefixOp) E() {}

type BinOp struct {
	Lhs      Expr
	Operator string
	Rhs      Expr
}

func (b BinOp) E() {}
func (b *BinOp) String() string {
	str := fmt.Sprintf("%s %s %s", b.Lhs, b.Operator, b.Rhs)
	return str
}

type PostfixOp struct {
	Lhs      Expr
	Operator string
}
