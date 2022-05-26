package main

import (
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
	scanner := newScanner(input)
	for {
		tok := scanner.next()
		if tok.typ == t_Eof {
			break
		}
		fmt.Println(tok)
	}
}
