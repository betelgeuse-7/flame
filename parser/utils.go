package parser

import (
	"errors"
	"flame/ast"
	"flame/token"
	"math"
	"strconv"
)

// TODO refactor this
//
// return parsed value from valLiteral, according to the data type (dt).
// returned ast.Expr value can be used to set *ast.VarDeclstm.Value, for example.
func giveProperValue(dt token.TokenType, valLiteral string) (ast.Expr, error) {
	switch dt {
	case token.T_StringKw:
		return &ast.StringLiteral{Val: valLiteral}, nil
	case token.T_IntKw, token.T_Int32Kw:
		val, err := strconv.ParseInt(valLiteral, 10, 64)
		if err != nil {
			return nil, err
		}
		if dt == token.T_Int32Kw {
			return &ast.I32Literal{ValStr: valLiteral, Val: int32(val)}, nil
		}
		return &ast.IntLiteral{ValStr: valLiteral, Val: val}, nil
	case token.T_UintKw, token.T_Uint32Kw:
		val, err := strconv.ParseUint(valLiteral, 10, 64)
		if err != nil {
			return nil, err
		}
		if dt == token.T_Uint32Kw {
			return &ast.U32Literal{ValStr: valLiteral, Val: uint32(val)}, nil
		}
		return &ast.UintLiteral{ValStr: valLiteral, Val: val}, nil
	case token.T_Float64Kw, token.T_Float32Kw:
		val, err := strconv.ParseFloat(valLiteral, 64)
		if err != nil {
			return nil, err
		}
		if dt == token.T_Float32Kw {
			return &ast.F32Literal{ValStr: valLiteral, Val: float32(val)}, nil
		}
		return &ast.FloatLiteral{ValStr: valLiteral, Val: val}, nil
	case token.T_BoolKw:
		val, err := strconv.ParseBool(valLiteral)
		if err != nil {
			return nil, err
		}
		return &ast.BooleanLiteral{ValStr: valLiteral, Val: val}, nil
	}
	return nil, errors.New("not a primitive value")
}

func checkPrimitiveValue(p *Parser, dt token.TokenType) bool {
	switch dt {
	case token.T_StringKw:
		if p.cur.Typ != token.T_String {
			p.reportErr("%d:%d Expected next token to be %s, got %s instead", p.cur.Pos.Y, p.cur.Pos.X, token.T_String, p.peek.Typ)
			return false
		}
	case token.T_UintKw, token.T_Uint32Kw:
		if ok := checkIsUint(dt, p.cur.Lit); !ok {
			p.reportErr("%d:%d invalid uint/u32 value: '%s'", p.cur.Pos.Y, p.cur.Pos.X, p.cur.Lit)
			return false
		}
	case token.T_IntKw, token.T_Int32Kw:
		if ok := checkIsInt(dt, p.cur.Lit); !ok {
			p.reportErr("%d:%dinvalid int/i32 value: '%s'", p.cur.Pos.Y, p.cur.Pos.X, p.cur.Lit)
			return false
		}
	case token.T_BoolKw:
		if p.cur.Lit != "true" && p.cur.Lit != "false" {
			p.reportErr("%d:%d invalid value for bool type: '%s'", p.cur.Pos.Y, p.cur.Pos.X, p.cur.Lit)
			return false
		}
	case token.T_Float64Kw, token.T_Float32Kw:
		if ok := checkIsFloat(dt, p.cur.Lit); !ok {
			p.reportErr("%d:%d invalid float/f32 value: '%s'", p.cur.Pos.Y, p.cur.Pos.X, p.cur.Lit)
			return false
		}
	default:
		panic("parseConstDecl: '" + string(dt) + "' >>> NOT IMPLEMENTED")
	}
	return true
}

// ? These functions look very alike. vvvv

func checkIsUint(dt token.TokenType, val string) bool {
	if dt == token.T_Uint32Kw {
		var maxU32Val uint32 = math.MaxUint32
		u32Val, err := strconv.ParseUint(val, 10, 32)
		if err != nil {
			return false
		}
		if u32Val > uint64(maxU32Val) {
			return false
		}
		return true
	}
	// dt is token.T_UintKw
	var maxU64Val uint64 = math.MaxUint64
	u64Val, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return false
	}
	if u64Val > maxU64Val {
		return false
	}
	return true
}

func checkIsInt(dt token.TokenType, val string) bool {
	if dt == token.T_Int32Kw {
		var maxI32Val int32 = math.MaxInt32
		i32Val, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return false
		}
		if i32Val > int64(maxI32Val) {
			return false
		}
		return true
	}
	var maxI64Val int64 = math.MaxInt64
	i64Val, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return false
	}
	if i64Val > maxI64Val {
		return false
	}
	return true
}

func checkIsFloat(dt token.TokenType, val string) bool {
	if dt == token.T_Float32Kw {
		var maxF32Val float32 = math.MaxFloat32
		f32Val, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return false
		}
		if f32Val > float64(maxF32Val) {
			return false
		}
		return true
	}
	var maxF64Val float64 = math.MaxFloat64
	f64Val, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return false
	}
	if f64Val > maxF64Val {
		return false
	}
	return true
}

func isDataTypeKw(typ token.TokenType) bool {
	for _, v := range dataTypeKws {
		if typ == v {
			return true
		}
	}
	return false
}
