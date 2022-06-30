package parser

import (
	"flame/ast"
	"flame/scanner"
	"flame/token"
	"strconv"
)

// TODO Fix line numbers in parser errors.

type Parser struct {
	scanner *scanner.Scanner
	cur     token.Token
	peek    token.Token
	errors  []string

	// infix expression parsing methods
	infixParseFns map[token.TokenType]infixParseFn
}

func New(scanner *scanner.Scanner) *Parser {
	p := &Parser{
		scanner:       scanner,
		infixParseFns: make(map[token.TokenType]infixParseFn),
	}
	p.registerAllExprParseFns()
	p.advance()
	p.advance()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
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
	if p.cur.Typ == token.T_Eof {
		return nil
	}
	switch p.cur.Typ {
	case token.T_Octothorp:
		return p.parseConstDecl()
	case token.T_If:
		return p.parseIfStmt()
	default:
		return p.parseExprStmt()
	}
}

func (p *Parser) parseVarDecl() *ast.VarDeclStmt {
	s := &ast.VarDeclStmt{DataType: p.cur.Typ}
	if ok := p.assertPeek(token.T_Ident); !ok {
		return nil
	}
	s.Name = p.cur.Lit
	if ok := p.assertPeek(token.T_Eq); !ok {
		return nil
	}
	p.advance()
	if ok := checkPrimitiveValue(p, s.DataType); !ok {
		return nil
	}
	s.Value = p.parseExpr()
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
	if ok := p.assertPeek(token.T_Ident); !ok {
		return nil
	}
	s.Decl.Name = p.cur.Lit
	if ok := p.assertPeek(token.T_Eq); !ok {
		return nil
	}
	p.advance()
	if ok := checkPrimitiveValue(p, s.Decl.DataType); !ok {
		return nil
	}
	s.Decl.Value = p.parseExpr()
	return s
}

func (p *Parser) parseExprStmt() *ast.ExprStmt {
	stmt := &ast.ExprStmt{}
	stmt.Expr = p.parseExpr()
	return stmt
}

func (p *Parser) parseExpr() ast.Expr {
	// TODO error handling
	var leftExpr ast.Expr
	switch p.cur.Typ {
	case token.T_Uint, token.T_Uint32:
		leftExpr = p.parseUnsignedIntegerLiteral()
	case token.T_Int, token.T_Int32:
		leftExpr = p.parseSignedIntegerLiteral()
	case token.T_Float32, token.T_Float64:
		leftExpr = p.parseFloatLiteral()
	case token.T_Bool:
		leftExpr = p.parseBoolLiteral()
	case token.T_String:
		leftExpr = p.parseStringLiteral()
	}
	// is this an infix expr, or a prefix expr ?
	infixFn := p.infixParseFns[p.peek.Typ]
	if infixFn == nil {
		return leftExpr
	}
	p.advance()
	expr := infixFn(leftExpr)
	return expr
}

func (p *Parser) parseStringLiteral() *ast.StringLiteral {
	return &ast.StringLiteral{Val: p.cur.Lit}
}

func (p *Parser) parseSignedIntegerLiteral() ast.SignedIntegerLiteral {
	switch p.cur.Typ {
	case token.T_Int:
		val, err := strconv.ParseInt(p.cur.Lit, 10, 64)
		if err != nil {
			return nil
		}
		return &ast.IntLiteral{ValStr: p.cur.Lit, Val: val}
	case token.T_Int32:
		val, err := strconv.ParseInt(p.cur.Lit, 10, 32)
		if err != nil {
			return nil
		}
		return &ast.I32Literal{ValStr: p.cur.Lit, Val: int32(val)}
	}
	return nil
}

func (p *Parser) parseUnsignedIntegerLiteral() ast.UnsignedIntegerLiteral {
	switch p.cur.Typ {
	case token.T_Uint:
		val, err := strconv.ParseUint(p.cur.Lit, 10, 64)
		if err != nil {
			return nil
		}
		return &ast.UintLiteral{ValStr: p.cur.Lit, Val: val}
	case token.T_Uint32:
		val, err := strconv.ParseUint(p.cur.Lit, 10, 32)
		if err != nil {
			return nil
		}
		return &ast.U32Literal{ValStr: p.cur.Lit, Val: uint32(val)}
	}
	return nil
}

func (p *Parser) parseFloatLiteral() ast.IFloatLiteral {
	switch p.cur.Typ {
	case token.T_Float64:
		val, err := strconv.ParseFloat(p.cur.Lit, 64)
		if err != nil {
			return nil
		}
		return &ast.FloatLiteral{ValStr: p.cur.Lit, Val: val}
	case token.T_Float32:
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

func (p *Parser) parseInfixExpr(left ast.Expr) ast.Expr {
	expr := &ast.BinOp{
		Lhs:      left,
		Operator: p.cur.Lit,
	}
	p.advance()
	expr.Rhs = p.parseExpr()
	if expr.Rhs == nil {
		x := p.cur.Pos.X
		y := p.cur.Pos.Y
		p.reportErr("%d:%d Expected another expression after '%s', but got nothing", y, x, expr.Operator)
		return nil
	}
	return expr
}

// TODO Re-write if statement parsing method below vvv

/* // TODO how can i simplify, and make the code below, maintainable?
// ? This is so ugly
func (p *Parser) parseIfStmt() *ast.IfStmt {
	stmt := &ast.IfStmt{}
	p.advance()
	stmt.Cond = p.parseExpr()
	if stmt.Cond == nil {
		p.reportErr("%d:%d No condition in if statement", p.cur.Pos.Y, p.cur.Pos.X)
		return nil
	}
	if ok := p.assertPeek(token.T_LCurly); !ok {
		return nil
	}
	p.advance()
	for p.cur.Typ != token.T_RCurly {
		if p.cur.Typ == token.T_Eof {
			break
		}
		if bodyStmt := p.parseStmt(); bodyStmt != nil {
			stmt.Body = append(stmt.Body, bodyStmt)
		}
		p.advance()
	}
	if ok := p.assertCur(token.T_RCurly); !ok {
		return nil
	}
	if p.cur.Typ == token.T_ElseIf {
		stmt.Alternative = p.parseIfStmt()
	}
	fmt.Println("parseif: ", p.cur.Typ, p.cur.Lit)
	if p.cur.Typ == token.T_Else {
		elseBranch := []ast.Stmt{}
		if ok := p.assertPeek(token.T_LCurly); !ok {
			return nil
		}
		for p.cur.Typ != token.T_RCurly {
			if p.cur.Typ == token.T_Eof {
				break
			}
			if elseBranchStmt := p.parseStmt(); elseBranchStmt != nil {
				elseBranch = append(elseBranch, elseBranchStmt)
			}
			p.advance()
		}
		if ok := p.assertCur(token.T_RCurly); !ok {
			return nil
		}
		stmt.Default = elseBranch
		p.advance()
	}
	return stmt
}
*/
