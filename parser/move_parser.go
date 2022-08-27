package parser

import (
	"flame/token"
	"fmt"
)

func (p *Parser) advance() {
	p.cur = p.peek
	p.peek = p.scanner.Next()
}

func (p *Parser) reportErr(format string, args ...interface{}) {
	err_ := fmt.Sprintf(format, args...)
	lineNos := fmt.Sprintf("%d:%d ", p.cur.Pos.Y, p.cur.Pos.X)
	err_ = lineNos + err_
	p.errors = append(p.errors, err_)
}

// expect next (peek) token to be tok
// call p.advance if not disappointed :P
func (p *Parser) expect(tok token.TokenType) bool {
	msg := "expected next token to be %s, got %s instead"
	illegalMsg := "expected next token to be %s, got %s ('%s') instead"
	if p.peek.Typ != tok {
		if p.peek.Typ == token.Illegal {
			p.reportErr(illegalMsg, tok, p.peek.Typ, p.peek.Lit)
		} else {
			p.reportErr(msg, tok, p.peek.Typ)
		}
		return false
	}
	p.advance()
	return true
}

func (p *Parser) expectType(isConst bool) bool {
	msg := "expected a type"
	if isConst {
		msg += " after '#'"
	}
	peek := p.peek.Typ
	if peek == token.Illegal || peek == token.Newline || peek == token.Eof {
		p.reportErr(msg)
		return false
	}
	p.advance()
	return true
}
