package main

import (
	"flame/parser"
	"flame/scanner"
	"fmt"
	"log"
	"os"
)

// scanner --> parser --> env      --> typecheck --> codegen
// tokens  --> ast    --> env,ast  --> ast       --> code
//
// eval, evaluates the ast, and creates an environment, to keep the variables, types, etc. in.
// typecheck, checks if there's any type errors (e.g adding an int and a string).
// codegen spits out code. (Go)

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
	if errs := p.Errors(); len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e)
		}
	}
	fmt.Println(program.Stmts)
	/*
		declarations := gogen.Compile(program)
		goCode := "package main\n\nfunc main() {\n"
		goCode += declarations
		goCode += "}"
		if err := os.WriteFile(fileName+"_compiled.go.example", []byte(goCode), 0777); err != nil {
			log.Fatalln("error while writing to file: ", err.Error())
		}
		log.Println("Compiled Flame to Go :)")
	*/
}
