package main

import "flame/token"

type Parser struct {
	Toks   []token.Token
	Pos    int
	Tok    token.Token
	errors []string
}

func New(tokenStream []token.Token) *Parser {
	p := &Parser{Toks: tokenStream, Pos: 0}
	p.Tok = p.Toks[p.Pos]
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}
