package main

import (
	"flame/scanner"
	"flame/token"
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
	for {
		tok := scanner.Next()
		if tok.Typ == token.T_Eof {
			break
		}
		fmt.Println(tok)
	}
}
