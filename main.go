package main

import (
	"flame/ast"
	"flame/codegen/gogen"
	"flame/parser"
	"flame/scanner"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("I need a file name")
	}
	fileName := os.Args[1]
	lenFn := len(fileName)
	if fileName[lenFn-3:lenFn] != ".fl" {
		log.Fatalln("A file with .fl file extension is needed")
	}
	bx, err := os.ReadFile(fileName)
	if err != nil {
		panic("error while reading " + fileName + ": " + err.Error() + "\n")
	}
	input := string(bx)
	scanner := scanner.New(input)
	p := parser.New(scanner)
	program := p.ParseProgram()
	goCode := "package main\n\n"
	goCode += "func main() {\n\t"
	goCode += gogen.CompileVarDecl(program.Stmts[0].(*ast.VariableDeclarationStmt))
	goCode += gogen.CompileVarDecl(program.Stmts[1].(*ast.VariableDeclarationStmt))
	goCode += "}"
	//os.WriteFile(fileName+"_compiled.go", []byte(goCode), 0777)
	fmt.Println(goCode)
	log.Println("Compiled Flame to Go :)")
}
