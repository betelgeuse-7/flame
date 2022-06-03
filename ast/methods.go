package ast

import (
	"bytes"
	"fmt"
)

/********* *ast.Program ************************/

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Stmts {
		out.WriteString(s.String())
	}
	return out.String()
}

/*************************************************/

/********** *ast.VariableDeclarationStmt **********/

func (v *VariableDeclarationStmt) String() string {
	var out bytes.Buffer
	//repr := fmt.Sprintf("%s %s = %s", v.Tok.Lit, v.Name, v.Value.String())
	repr := fmt.Sprintf("%s %s = %s", v.DataType.Lit, v.Name, v.Value)
	out.WriteString(repr)
	return out.String()
}

func (v *VariableDeclarationStmt) stmtNode() {}

/***************************************************/
/******* *ast.Identifier ****************************/
/*
func (i *Identifier) expressionNode()      {}
func (i *Identifier) String() string       { return i.Value }
*/
/****************************************************/
