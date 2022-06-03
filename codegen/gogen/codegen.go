package gogen

import (
	"flame/ast"
	"fmt"
)

// temporary function
// this will die
func CompileVarDecl(d *ast.VariableDeclarationStmt) string {
	dataType := d.DataType
	name := d.Name
	val := d.Value
	return fmt.Sprintf("var %s %s = %s\n\t", name, dataType.Lit, val)
}
