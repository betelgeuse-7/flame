package main

type parser struct {
	toks   []token
	pos    int
	tok    token
	errors []string
}

func newParser(tokenStream []token) *parser {
	p := &parser{toks: tokenStream, pos: 0}
	p.tok = p.toks[p.pos]
	return p
}
