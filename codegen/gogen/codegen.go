package gogen

var flameGoKws = map[string]string{
	"string": "string",
	"uint":   "uint64",
	"u32":    "uint32",
	"f32":    "float32",
	"float":  "float64",
	"int":    "int64",
	"i32":    "int32",
	"bool":   "bool",
}

/*
func Compile(program *ast.Program) string {
	res := ""
	for _, stmt := range program.Stmts {
		switch stmt.(type) {
		case *ast.ConstDeclStmt:
			decl := stmt.(*ast.ConstDeclStmt)
			goDataType := flameGoKws[string(decl.Decl.DataType)]
			if decl.Decl.DataType == token.T_StringKw {
				res += fmt.Sprintf("var %s %s = \"%s\"\n", decl.Decl.Name, goDataType, decl.Decl.Value.Value())
			} else {
				res += fmt.Sprintf("var %s %s = %s\n", decl.Decl.Name, goDataType, decl.Decl.Value.Value())

			}
		case *ast.VarDeclStmt:
			decl := stmt.(*ast.VarDeclStmt)
			goDataType := flameGoKws[string(decl.DataType)]
			if decl.DataType == token.T_StringKw {
				res += fmt.Sprintf("var %s %s = \"%s\"\n", decl.Name, goDataType, decl.Value)
			} else {
				res += fmt.Sprintf("var %s %s = %s\n", decl.Name, goDataType, decl.Value)
			}
		}
	}
	return res
}
*/
