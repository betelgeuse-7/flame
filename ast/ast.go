package ast

import (
	"flame/token"
	"fmt"
)

// TODO ast.Expr interface ???

type Program struct {
	Stmts []Stmt
}

type Node interface{}

type Stmt interface {
	Node
	S()
}

type Expr interface {
	Node
	E()
	Value() string
}

type VarDeclStmt struct {
	DataType token.TokenType
	Name     string
	Value    Expr
}

func (v VarDeclStmt) S() {}
func (v *VarDeclStmt) String() string {
	return fmt.Sprintf("%s %s = %s", string(v.DataType), v.Name, v.Value.Value())
}

type ConstDeclStmt struct {
	Decl      VarDeclStmt
	Octothorp token.TokenType
}

func (c ConstDeclStmt) S() {}
func (c *ConstDeclStmt) String() string {
	return fmt.Sprintf("#%s %s = %s", string(c.Decl.DataType), c.Decl.Name, c.Decl.Value.Value())
}

/*
type PrimitiveValue struct {
	DataType token.TokenType
	Val      string
}

func (p PrimitiveValue) E()            {}
func (p PrimitiveValue) Value() string { return p.Val }
*/
