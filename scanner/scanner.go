package scanner

import (
	"flame/token"
	"strings"
)

const _EOF_RUNE = rune(0)

type Scanner struct {
	input []rune
	ch    rune
	pos   int
	x     int // col
	y     int // row
}

func New(input string) *Scanner {
	s := &Scanner{
		input: []rune(input),
		pos:   0,
		y:     1,
		x:     1,
	}
	s.ch = s.input[s.pos]
	return s
}

func (s *Scanner) advance() {
	if p := s.peek(); p == '\n' {
		s.y++
		s.x = 1
	} else {
		s.x++
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
		return token.Token{Typ: token.T_Eof, Lit: "EOF", Pos: token.TokenPos{X: s.x, Y: s.y}}
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
		doubleCharTok.Pos = token.TokenPos{X: s.x, Y: s.y}
		return doubleCharTok
	}
	singleCharTok, ok := singleCharMap[byte(s.ch)]
	// is it a number starting with a minus? (e.g -5)
	if singleCharTok.Typ == token.T_Minus {
		if peek := s.peek(); isDigit(peek) {
			return s.scanNumber()
		}
	}
	if ok {
		s.advance()
		singleCharTok.Pos = token.TokenPos{X: s.x, Y: s.y}
		return singleCharTok
	}
	illegalCh := s.ch
	s.advance()
	return token.Token{Typ: token.T_Illegal, Lit: string(illegalCh), Pos: token.TokenPos{X: s.x, Y: s.y}}
}

func (s *Scanner) scanNumber() token.Token {
	startsWithMinus := s.ch == '-'
	dots := 0
	start := s.pos
	s.advance()
	for isDigit(s.ch) || s.ch == '.' {
		// floats can't have more than one dot
		if dots > 1 {
			break
		}
		if s.ch == '.' {
			dots++
		}
		s.advance()
	}
	lit := string(s.input[start:s.pos])
	tok := token.Token{Lit: lit, Pos: token.TokenPos{X: s.x, Y: s.y}}
	if dots > 0 {
		tok.Typ = token.T_Float64
	} else if startsWithMinus {
		tok.Typ = token.T_Int
	} else {
		tok.Typ = token.T_Uint
	}
	return tok
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
		tok = token.Token{Typ: token.T_Ident, Lit: lit, Pos: token.TokenPos{X: s.x, Y: s.y}}
	}
	if tok.Typ != token.T_Ident {
		tok.Lit = strings.ToLower(string(tok.Typ))
	}
	if tok.Typ == token.T_Bool {
		tok.Lit = lit
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
	return token.Token{Typ: token.T_String, Lit: lit, Pos: token.TokenPos{X: s.x, Y: s.y}}
}

func (s *Scanner) scanComment() token.Token {
	// get rid of /
	s.advance()
	start := s.pos
	for s.ch != '\n' && s.ch != '\r' {
		s.advance()
	}
	lit := string(s.input[start:s.pos])
	return token.Token{Typ: token.T_Comment, Lit: lit, Pos: token.TokenPos{X: s.x, Y: s.y}}
}

func (s *Scanner) eatWs() {
	for isWhitespace(s.ch) {
		s.advance()
	}
}
