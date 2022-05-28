package scanner

import "flame/token"

type Scanner struct {
	input string
	ch    byte
	pos   int
}

func New(input string) *Scanner {
	s := &Scanner{
		input: input,
		pos:   0,
	}
	s.ch = s.input[s.pos]
	return s
}

func (s *Scanner) advance() {
	s.pos++
	if s.pos == len(s.input) {
		s.ch = 0
		return
	}
	s.ch = s.input[s.pos]
}

func (s *Scanner) Next() token.Token {
	if s.ch == 0 {
		return token.Token{Typ: token.T_Eof, Lit: "EOF"}
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
		return token.Token{Typ: token.T_Plus, Lit: "+"}
	case '-':
		s.advance()
		return token.Token{Typ: token.T_Minus, Lit: "-"}
	case '*':
		s.advance()
		return token.Token{Typ: token.T_Mul, Lit: "*"}
	case '/':
		s.advance()
		return token.Token{Typ: token.T_Div, Lit: "/"}
	case '(':
		s.advance()
		return token.Token{Typ: token.T_Lparen, Lit: "("}
	case ')':
		s.advance()
		return token.Token{Typ: token.T_Rparen, Lit: ")"}
	case '=':
		s.advance()
		return token.Token{Typ: token.T_Eq, Lit: "="}
	default:
		s.advance()
		return token.Token{Typ: token.T_Illegal, Lit: string(s.ch)}
	}
}

func (s *Scanner) scanWs() token.Token {
	start := s.pos
	for isWhitespace(s.ch) {
		s.advance()
	}
	lit := s.input[start:s.pos]
	return token.Token{Typ: token.T_Whitespace, Lit: lit}
}

func (s *Scanner) scanNumber() token.Token {
	start := s.pos
	for isDigit(s.ch) || s.ch == '.' {
		s.advance()
	}
	lit := s.input[start:s.pos]
	return token.Token{Typ: token.T_Number, Lit: lit}
}

func (s *Scanner) scanIdentOrKw() token.Token {
	start := s.pos
	for isAsciiLetter(s.ch) {
		s.advance()
	}
	lit := s.input[start:s.pos]
	tok := token.Token{}
	switch lit {
	case "var":
		tok = token.Token{Typ: token.T_Var}
	case "const":
		tok = token.Token{Typ: token.T_Const}
	case "string":
		tok = token.Token{Typ: token.T_StringKw}
	case "uint":
		tok = token.Token{Typ: token.T_UintKw}
	case "int":
		tok = token.Token{Typ: token.T_IntKw}
	case "bool":
		tok = token.Token{Typ: token.T_BoolKw}
	case "println":
		tok = token.Token{Typ: token.T_Println}
	default:
		tok = token.Token{Typ: token.T_Ident, Lit: lit}
	}
	if tok.Typ != token.T_Ident {
		tok.Lit = string(tok.Typ)
	}
	return tok
}

func (s *Scanner) scanString() token.Token {
	// do not include quotes in Token.Lit
	s.advance()
	start := s.pos
	for s.ch != '"' && s.ch != 0 {
		s.advance()
	}
	lit := s.input[start:s.pos]
	// skip over the terminating quote
	s.advance()
	return token.Token{Typ: token.T_String, Lit: lit}
}
