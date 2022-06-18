package ast

import (
	"flame/token"
	"fmt"
)

type Program struct {
	Stmts []Stmt
}

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
	IsANumericValue() bool
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
	Operator string
	Rhs      Expr
}

func (p PrefixOp) E()                    {}
func (p PrefixOp) IsANumericValue() bool { return false }

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
func (b BinOp) IsANumericValue() bool { return false }

type PostfixOp struct {
	Lhs      Expr
	Operator string
}

func (p PostfixOp) IsANumericValue() bool { return false }

type IfStmt struct {
	Cond        Expr
	Body        []Stmt
	Alternative *IfStmt // elseif
	Default     []Stmt  // else
}

func (i *IfStmt) S() {}
func (i *IfStmt) String() string {
	res := fmt.Sprintf("if %s {\n", i.Cond.String())
	if len(i.Body) > 0 {
		for _, v := range i.Body {
			res += "\t" + v.String() + "\n"
		}
	}
	if i.Alternative != nil {
		res += "} else"
		res += i.Alternative.String()
	}
	if len(i.Default) > 0 {
		res += " else {\n\t"
		for _, v := range i.Default {
			res += v.String() + "\n"
		}
	}
	res += "}"
	return res
}
