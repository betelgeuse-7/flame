package ast

import (
	"flame/token"
	"fmt"
)

type Program struct {
	Stmts []Stmt
}

type Node interface{}

type Stmt interface {
	Node
	S()
	String() string
}

type Expr interface {
	Node
	E()
	String() string
}

type ExprStmt struct {
	Expr
}

func (e ExprStmt) S() {}

type VarDeclStmt struct {
	DataType token.TokenType
	Name     string
	Value    Expr
}

func (v VarDeclStmt) S() {}
func (v *VarDeclStmt) String() string {
	return fmt.Sprintf("%s %s = %s", string(v.DataType), v.Name, v.Value)
}

type ConstDeclStmt struct {
	Decl      VarDeclStmt
	Octothorp token.TokenType
}

func (c ConstDeclStmt) S() {}
func (c *ConstDeclStmt) String() string {
	return fmt.Sprintf("#%s %s = %s", string(c.Decl.DataType), c.Decl.Name, c.Decl.Value)
}

type PrefixOp struct {
	Operator
	Rhs Expr
}

func (p PrefixOp) E() {}

type BinOp struct {
	Lhs Expr
	Operator
	Rhs Expr
}

func (b BinOp) E() {}

type PostfixOp struct {
	Lhs Expr
	Operator
}
