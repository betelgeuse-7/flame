package ast

import (
	"flame/token"
)

/// TODO Position information for each node.

type Node interface{}

type Stmt interface {
	Node
	stmtNode()
}

type Expr interface {
	Node
	exprNode()
}

type Program struct {
	Stmts []Stmt
}

type VariableDeclarationStmt struct {
	DataType token.Token // data type kw token
	Name     string      // variable name
	Value    string
}
