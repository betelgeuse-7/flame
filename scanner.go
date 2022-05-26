package main

type scanner struct {
	input string
	ch    byte
	pos   int
}

func newScanner(input string) *scanner {
	s := &scanner{
		input: input,
		pos:   0,
	}
	s.ch = s.input[s.pos]
	return s
}

func (s *scanner) advance() {
	s.pos++
	if s.pos == len(s.input) {
		s.ch = 0
		return
	}
	s.ch = s.input[s.pos]
}

func (s *scanner) next() token {
	if s.ch == 0 {
		return token{typ: t_Eof, lit: "EOF"}
	}
	if isWhitespace(s.ch) {
		return s.scanWs()
	} else if isDigit(s.ch) {
		return s.scanNumber()
	} else if isAsciiLetter(s.ch) {
		return s.scanIdentOrKw()
	} else if s.ch == '"' {
		return s.scanString()
	}
	switch s.ch {
	case '+':
		s.advance()
		return token{typ: t_Plus, lit: "+"}
	case '-':
		s.advance()
		return token{typ: t_Minus, lit: "-"}
	case '*':
		s.advance()
		return token{typ: t_Mul, lit: "*"}
	case '/':
		s.advance()
		return token{typ: t_Div, lit: "/"}
	case '(':
		s.advance()
		return token{typ: t_Lparen, lit: "("}
	case ')':
		s.advance()
		return token{typ: t_Rparen, lit: ")"}
	case '=':
		s.advance()
		return token{typ: t_Eq, lit: "="}
	default:
		s.advance()
		return token{typ: t_Illegal, lit: string(s.ch)}
	}
}

func (s *scanner) scanWs() token {
	start := s.pos
	for isWhitespace(s.ch) {
		s.advance()
	}
	lit := s.input[start:s.pos]
	return token{typ: t_Whitespace, lit: lit}
}

func (s *scanner) scanNumber() token {
	start := s.pos
	for isDigit(s.ch) || s.ch == '.' {
		s.advance()
	}
	lit := s.input[start:s.pos]
	return token{typ: t_Number, lit: lit}
}

func (s *scanner) scanIdentOrKw() token {
	start := s.pos
	for isAsciiLetter(s.ch) {
		s.advance()
	}
	lit := s.input[start:s.pos]
	tok := token{}
	switch lit {
	case "var":
		tok = token{typ: t_Var}
	case "const":
		tok = token{typ: t_Const}
	case "string":
		tok = token{typ: t_StringKw}
	case "uint":
		tok = token{typ: t_UintKw}
	case "int":
		tok = token{typ: t_IntKw}
	case "bool":
		tok = token{typ: t_BoolKw}
	case "println":
		tok = token{typ: t_Println}
	default:
		tok = token{typ: t_Ident, lit: lit}
	}
	if tok.typ != t_Ident {
		tok.lit = string(tok.typ)
	}
	return tok
}

func (s *scanner) scanString() token {
	// do not include quotes in token.lit
	s.advance()
	start := s.pos
	for s.ch != '"' && s.ch != 0 {
		s.advance()
	}
	lit := s.input[start:s.pos]
	// skip over the terminating quote
	s.advance()
	return token{typ: t_String, lit: lit}
}
