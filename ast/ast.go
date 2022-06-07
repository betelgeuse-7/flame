package ast

import "flame/token"

type Program struct {
	Stmts []Stmt
}

type Stmt interface{ S() }
type Expr interface{ E() }

type VarDeclStmt struct {
	DataType token.TokenType
	Name     string
	Value    Expr
}

func (v VarDeclStmt) S() {}

type ConstDeclStmt struct {
	Decl      VarDeclStmt
	Octothorp token.TokenType
}

func (c ConstDeclStmt) S() {}
