package scanner

import (
	"flame/token"
)

const _EOF_RUNE = rune(0)

type Scanner struct {
	input []rune
	ch    rune
	pos   int // col
	y     int // row
}

func New(input string) *Scanner {
	s := &Scanner{
		input: []rune(input),
		pos:   0,
		y:     1,
	}
	s.ch = s.input[s.pos]
	return s
}

func (s *Scanner) advance() {
	if p := s.peek(); p == '\n' {
		s.y++
	}
	s.pos++
	if s.pos == len(s.input) {
		s.ch = _EOF_RUNE
		return
	}
	s.ch = s.input[s.pos]
}

func (s *Scanner) doubleAdvance() {
	s.advance()
	s.advance()
}

func (s *Scanner) peek() rune {
	if s.pos+1 == len(s.input) {
		return _EOF_RUNE
	}
	return s.input[s.pos+1]
}

func (s *Scanner) Next() token.Token {
	if s.ch == _EOF_RUNE {
		return token.Token{Typ: token.T_Eof, Lit: "EOF"}
	}
	if isWhitespace(s.ch) {
		s.eatWs()
		return s.Next()
	} else if isDigit(s.ch) {
		return s.scanNumber()
	} else if isAsciiLetter(s.ch) {
		return s.scanIdentOrKw()
	} else if s.ch == '"' {
		return s.scanString()
	}
	ahead := s.peek()
	curAndAhead := string(s.ch) + string(ahead)
	if curAndAhead == "//" {
		s.advance()
		return s.scanComment()
	}
	doubleCharTok, ok := doubleCharMap[curAndAhead]
	if ok {
		s.doubleAdvance()
		return doubleCharTok
	}
	singleCharTok, ok := singleCharMap[byte(s.ch)]
	if ok {
		s.advance()
		return singleCharTok
	}
	illegalCh := s.ch
	s.advance()
	return token.Token{Typ: token.T_Illegal, Lit: string(illegalCh)}
}

func (s *Scanner) scanNumber() token.Token {
	start := s.pos
	for isDigit(s.ch) || s.ch == '.' {
		s.advance()
	}
	lit := string(s.input[start:s.pos])
	return token.Token{Typ: token.T_Number, Lit: lit}
}

func (s *Scanner) scanIdentOrKw() token.Token {
	start := s.pos
	for isAsciiLetter(s.ch) || isDigit(s.ch) {
		s.advance()
	}
	lit := string(s.input[start:s.pos])
	tok := token.Token{}
	tok, ok := keywordMap[lit]
	if !ok {
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
	lit := string(s.input[start:s.pos])
	// skip over the terminating quote
	s.advance()
	return token.Token{Typ: token.T_String, Lit: lit}
}

func (s *Scanner) scanComment() token.Token {
	// get rid of /
	s.advance()
	start := s.pos
	for s.ch != '\n' && s.ch != '\r' {
		s.advance()
	}
	lit := string(s.input[start:s.pos])
	return token.Token{Typ: token.T_Comment, Lit: lit}
}

func (s *Scanner) eatWs() {
	for isWhitespace(s.ch) {
		s.advance()
	}
}
