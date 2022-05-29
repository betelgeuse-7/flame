package scanner

import (
	"flame/token"
)

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

func (s *Scanner) doubleAdvance() {
	s.advance()
	s.advance()
}

func (s *Scanner) peek() byte {
	if s.pos+1 == len(s.input) {
		// eof
		return 0
	}
	return s.input[s.pos+1]
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
	singleCharTok, ok := singleCharMap[s.ch]
	if ok {
		s.advance()
		return singleCharTok
	}
	return token.Token{Typ: token.T_Illegal, Lit: string(s.ch)}
	/*
		switch s.ch {
		case '+':
			if p := s.peek(); p == '+' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_PlusPlus, Lit: "++"}
			} else if p == '=' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_PlusEq, Lit: "+="}
			}
			s.advance()
			return token.Token{Typ: token.T_Plus, Lit: "+"}
		case '-':
			if p := s.peek(); p == '-' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_MinusMinus, Lit: "--"}
			} else if p == '=' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_MinusEq, Lit: "-="}
			} else if p == '>' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_SingleArrow, Lit: "->"}
			}
			s.advance()
			return token.Token{Typ: token.T_Minus, Lit: "-"}
		case '*':
			if p := s.peek(); p == '=' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_MulEq, Lit: "*="}
			}
			s.advance()
			return token.Token{Typ: token.T_Mul, Lit: "*"}
		case '/':
			if p := s.peek(); p == '=' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_DivEq, Lit: "/="}
			} else if p == '/' {
				s.advance()
				return s.scanComment()
			}
			s.advance()
			return token.Token{Typ: token.T_Div, Lit: "/"}
		case '(':
			s.advance()
			return token.Token{Typ: token.T_Lparen, Lit: "("}
		case ')':
			s.advance()
			return token.Token{Typ: token.T_Rparen, Lit: ")"}
		case '=':
			if p := s.peek(); p == '=' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_DoubleEq, Lit: "=="}
			} else if p == '>' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_DoubleArrow, Lit: "=>"}
			}
			s.advance()
			return token.Token{Typ: token.T_Eq, Lit: "="}
		case '<':
			if p := s.peek(); p == '<' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_BitLeftShift, Lit: "<<"}
			} else if p == '=' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_LtEq, Lit: "<="}
			}
			s.advance()
			return token.Token{Typ: token.T_Lt, Lit: "<"}
		case '>':
			if p := s.peek(); p == '>' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_BitRightShift, Lit: ">>"}
			} else if p == '=' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_GtEq, Lit: ">="}
			}
			s.advance()
			return token.Token{Typ: token.T_Gt, Lit: ">"}
		case '&':
			if p := s.peek(); p == '&' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_And, Lit: "&&"}
			}
			s.advance()
			return token.Token{Typ: token.T_BitAnd, Lit: "&"}
		case '|':
			if p := s.peek(); p == '|' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_Or, Lit: "||"}
			}
			s.advance()
			return token.Token{Typ: token.T_BitOr, Lit: "|"}
		case '^':
			s.advance()
			return token.Token{Typ: token.T_BitXor, Lit: "^"}
		case '!':
			if p := s.peek(); p == '=' {
				s.doubleAdvance()
				return token.Token{Typ: token.T_NotEq, Lit: "!="}
			}
			s.advance()
			return token.Token{Typ: token.T_Exclamation, Lit: "!"}
		case '{':
			s.advance()
			return token.Token{Typ: token.T_LCurly, Lit: "{"}
		case '}':
			s.advance()
			return token.Token{Typ: token.T_RCurly, Lit: "}"}
		case ',':
			s.advance()
			return token.Token{Typ: token.T_Comma, Lit: ","}
		case '$':
			s.advance()
			return token.Token{Typ: token.T_Dollar, Lit: "$"}
		case '.':
			s.advance()
			return token.Token{Typ: token.T_Dot, Lit: "."}
		default:
			s.advance()
			return token.Token{Typ: token.T_Illegal, Lit: string(s.ch)}
		}
	*/
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
	lit := s.input[start:s.pos]
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
	lit := s.input[start:s.pos]
	return token.Token{Typ: token.T_Comment, Lit: lit}
}
