package parser

import (
	"flame/token"
	"fmt"
)

func (p *Parser) advance() {
	p.cur = p.peek
	p.peek = p.scanner.Next()
}

func (p *Parser) assertCur(tt token.TokenType) bool {
	if p.cur.Typ == tt {
		p.advance()
		return true
	}
	p.reportErr(fmt.Sprintf("wanted '%s', got (%s, %s)", tt, p.cur.Typ, p.cur.Lit))
	return false
}

func (p *Parser) assertPeek(tt token.TokenType) bool {
	if p.peek.Typ == tt {
		p.advance()
		return true
	}
	p.reportErr(fmt.Sprintf("expected '%s', got (%s, %s)", tt, p.peek.Typ, p.peek.Lit))
	return false
}

func (p *Parser) reportErr(format string, args ...interface{}) {
	err_ := fmt.Sprintf(format, args...)
	lineNos := fmt.Sprintf("%d:%d ", p.cur.Pos.Y, p.cur.Pos.X)
	err_ = lineNos + err_
	p.errors = append(p.errors, err_)
}
