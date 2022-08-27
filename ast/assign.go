package ast

import (
	"flame/token"
	"fmt"
	"strings"
)

type Ident struct {
	Pos   token.Pos
	Ident string
}

func (i Ident) String() string {
	return i.Ident
}
func (Ident) E() {}

// represents a const or a var decl.
type GenericDeclStmt struct {
	TokPos token.Pos
	Tok    token.TokenType // token.Var || token.Octothorp
	Type   Type
	Ident  Ident // variable/constant name
	Value  Expr
}

type VarDeclStmt struct {
	GenericDeclStmt
}

func (v VarDeclStmt) String() string {
	return fmt.Sprintf("%s %s = %s", strings.ToLower(string(v.Tok)), v.Ident, v.Value.String())
}
func (VarDeclStmt) S() {}

type ConstDeclStmt struct {
	GenericDeclStmt
}

func (c ConstDeclStmt) String() string {
	return fmt.Sprintf("#%s %s = %s", strings.ToLower(string(c.Tok)), c.Ident, c.Value.String())
}
func (ConstDeclStmt) S() {}
