package parser

import (
	"flame/ast"
	"flame/scanner"
	"flame/token"
	"fmt"
	"strconv"
)

// TODO Fix line numbers in parser errors.

type Parser struct {
	scanner *scanner.Scanner
	cur     token.Token
	peek    token.Token
	errors  []string
}

func New(scanner *scanner.Scanner) *Parser {
	p := &Parser{
		scanner: scanner,
	}
	p.advance()
	p.advance()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) advance() {
	p.cur = p.peek
	p.peek = p.scanner.Next()
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peek.Typ == t {
		p.advance()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("%d:%d Expected next token to be %s, got %s instead", p.cur.Pos.Y, p.cur.Pos.X, t, p.peek.Typ)
	p.errors = append(p.errors, msg)
}

func (p *Parser) reportErr(format string, args ...interface{}) {
	p.errors = append(p.errors, fmt.Sprintf(format, args...))
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.cur.Typ != token.T_Eof {
		stmt := p.parseStmt()
		if stmt != nil {
			program.Stmts = append(program.Stmts, stmt)
		}
		p.advance()
	}
	return program
}

func (p *Parser) parseStmt() ast.Stmt {
	if isDataTypeKw(p.cur.Typ) {
		return p.parseVarDecl()
	}
	switch p.cur.Typ {
	case token.T_Octothorp:
		return p.parseConstDecl()
	default:
		return p.parseExprStmt()
	}
}

func (p *Parser) parseVarDecl() *ast.VarDeclStmt {
	s := &ast.VarDeclStmt{DataType: p.cur.Typ}
	if ok := p.expectPeek(token.T_Ident); !ok {
		return nil
	}
	s.Name = p.cur.Lit
	if ok := p.expectPeek(token.T_Eq); !ok {
		return nil
	}
	p.advance()
	if ok := checkPrimitiveValue(p, s.DataType); !ok {
		return nil
	}
	switch s.DataType {
	case token.T_StringKw:
		s.Value = p.parseStringLiteral()
	case token.T_IntKw, token.T_Int32Kw:
		val := p.parseSignedIntegerLiteral(s.DataType)
		if _, ok := val.(*ast.IntLiteral); ok {
			s.Value = val.(*ast.IntLiteral)
		} else if _, ok := val.(*ast.I32Literal); ok {
			s.Value = val.(*ast.I32Literal)
		}
	case token.T_UintKw, token.T_Uint32Kw:
		val := p.parseUnsignedIntegerLiteral(s.DataType)
		if _, ok := val.(*ast.UintLiteral); ok {
			s.Value = val.(*ast.UintLiteral)
		} else if _, ok := val.(*ast.U32Literal); ok {
			s.Value = val.(*ast.U32Literal)
		}
	case token.T_Float64Kw, token.T_Float32Kw:
		val := p.parseFloatLiteral(s.DataType)
		if _, ok := val.(*ast.FloatLiteral); ok {
			s.Value = val.(*ast.FloatLiteral)
		} else if _, ok := val.(*ast.F32Literal); ok {
			s.Value = val.(*ast.F32Literal)
		}
	case token.T_BoolKw:
		s.Value = p.parseBoolLiteral()
	}
	return s
}

func (p *Parser) parseConstDecl() *ast.ConstDeclStmt {
	s := &ast.ConstDeclStmt{Octothorp: p.cur.Typ}
	if !(isDataTypeKw(p.peek.Typ)) {
		p.reportErr("%d:%d Malformed Constant Declaration: expected a data type keyword after an octothorp, got %s", p.cur.Pos.Y, p.cur.Pos.X, p.peek.Lit)
		return nil
	}
	p.advance()
	s.Decl.DataType = p.cur.Typ
	if ok := p.expectPeek(token.T_Ident); !ok {
		return nil
	}
	s.Decl.Name = p.cur.Lit
	if ok := p.expectPeek(token.T_Eq); !ok {
		return nil
	}
	p.advance()
	if ok := checkPrimitiveValue(p, s.Decl.DataType); !ok {
		return nil
	}
	switch s.Decl.DataType {
	case token.T_StringKw:
		s.Decl.Value = p.parseStringLiteral()
	case token.T_IntKw, token.T_Int32Kw:
		val := p.parseSignedIntegerLiteral(s.Decl.DataType)
		if _, ok := val.(*ast.IntLiteral); ok {
			s.Decl.Value = val.(*ast.IntLiteral)
		} else if _, ok := val.(*ast.I32Literal); ok {
			s.Decl.Value = val.(*ast.I32Literal)
		}
	case token.T_UintKw, token.T_Uint32Kw:
		val := p.parseUnsignedIntegerLiteral(s.Decl.DataType)
		if _, ok := val.(*ast.UintLiteral); ok {
			s.Decl.Value = val.(*ast.UintLiteral)
		} else if _, ok := val.(*ast.U32Literal); ok {
			s.Decl.Value = val.(*ast.U32Literal)
		}
	case token.T_Float64Kw, token.T_Float32Kw:
		val := p.parseFloatLiteral(s.Decl.DataType)
		if _, ok := val.(*ast.FloatLiteral); ok {
			s.Decl.Value = val.(*ast.FloatLiteral)
		} else if _, ok := val.(*ast.F32Literal); ok {
			s.Decl.Value = val.(*ast.F32Literal)
		}
	case token.T_BoolKw:
		s.Decl.Value = p.parseBoolLiteral()
	}
	return s
}

func (p *Parser) parseExprStmt() *ast.ExprStmt {
	// PREFIX: 	++ + -- - ( * & ...
	// INFIX: 	+ - / * % ...
	// POSTFIX: ++ -- ' ...
	return nil
}

func (p *Parser) parseStringLiteral() *ast.StringLiteral {
	return &ast.StringLiteral{Val: p.cur.Lit}
}

func (p *Parser) parseSignedIntegerLiteral(dt token.TokenType) ast.SignedIntegerLiteral {
	switch dt {
	case token.T_IntKw:
		val, err := strconv.ParseInt(p.cur.Lit, 10, 64)
		if err != nil {
			return nil
		}
		return &ast.IntLiteral{ValStr: p.cur.Lit, Val: val}
	case token.T_Int32Kw:
		val, err := strconv.ParseInt(p.cur.Lit, 10, 32)
		if err != nil {
			return nil
		}
		return &ast.I32Literal{ValStr: p.cur.Lit, Val: int32(val)}
	}
	return nil
}

func (p *Parser) parseUnsignedIntegerLiteral(dt token.TokenType) ast.UnsignedIntegerLiteral {
	switch dt {
	case token.T_UintKw:
		val, err := strconv.ParseUint(p.cur.Lit, 10, 64)
		if err != nil {
			return nil
		}
		return &ast.UintLiteral{ValStr: p.cur.Lit, Val: val}
	case token.T_Uint32Kw:
		val, err := strconv.ParseUint(p.cur.Lit, 10, 32)
		if err != nil {
			return nil
		}
		return &ast.U32Literal{ValStr: p.cur.Lit, Val: uint32(val)}
	}
	return nil
}

func (p *Parser) parseFloatLiteral(dt token.TokenType) ast.IFloatLiteral {
	switch dt {
	case token.T_Float64Kw:
		val, err := strconv.ParseFloat(p.cur.Lit, 64)
		if err != nil {
			return nil
		}
		return &ast.FloatLiteral{ValStr: p.cur.Lit, Val: val}
	case token.T_Float32Kw:
		val, err := strconv.ParseFloat(p.cur.Lit, 32)
		if err != nil {
			return nil
		}
		return &ast.F32Literal{ValStr: p.cur.Lit, Val: float32(val)}
	}
	return nil
}

func (p *Parser) parseBoolLiteral() *ast.BooleanLiteral {
	val, err := strconv.ParseBool(p.cur.Lit)
	if err != nil {
		return nil
	}
	return &ast.BooleanLiteral{ValStr: p.cur.Lit, Val: val}
}
