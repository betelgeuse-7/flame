package ast

import (
	"flame/token"
	"testing"
)

func TestIfStmtString(t *testing.T) {
	stmt := &IfStmt{Cond: ExprStmt{Expr: &BooleanLiteral{ValStr: "true", Val: true}},
		Body: []Stmt{&VarDeclStmt{DataType: token.T_IntKw, Name: "x", Value: &IntLiteral{ValStr: "5", Val: 5}}},
		Alternative: &IfStmt{Cond: ExprStmt{Expr: &BooleanLiteral{ValStr: "false", Val: false}},
			Body: []Stmt{&ConstDeclStmt{Decl: VarDeclStmt{DataType: token.T_StringKw, Name: "name", Value: &StringLiteral{Val: "Jennifer"}}}}},
		Default: []Stmt{&ConstDeclStmt{Decl: VarDeclStmt{DataType: token.T_BoolKw, Name: "isRaining", Value: &BooleanLiteral{ValStr: "false", Val: false}}}}}

	s := stmt.String()
	t.Logf("%s\n", s)
}
