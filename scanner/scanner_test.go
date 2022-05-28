package scanner

import (
	"flame/token"
	"testing"
)

var (
	//_ws           = func(Lit string) token.Token { return token.Token{Typ: token.T_Whitespace, Lit: Lit} }
	_ident        = func(Lit string) token.Token { return token.Token{Typ: token.T_Ident, Lit: Lit} }
	_stringkw     = token.Token{Typ: token.T_StringKw, Lit: "STRINGKW"}
	_eq           = token.Token{Typ: token.T_Eq, Lit: "="}
	_string       = func(Lit string) token.Token { return token.Token{Typ: token.T_String, Lit: Lit} }
	_const        = token.Token{Typ: token.T_Const, Lit: "CONST"}
	_uintKw       = token.Token{Typ: token.T_UintKw, Lit: "UINTKW"}
	_number       = func(Lit string) token.Token { return token.Token{Typ: token.T_Number, Lit: Lit} }
	_println      = token.Token{Typ: token.T_PrintlnFn, Lit: "PRINTLNFN"}
	_lparen       = token.Token{Typ: token.T_Lparen, Lit: "("}
	_rparen       = token.Token{Typ: token.T_Rparen, Lit: ")"}
	_plus         = token.Token{Typ: token.T_Plus, Lit: "+"}
	_minus        = token.Token{Typ: token.T_Minus, Lit: "-"}
	_mul          = token.Token{Typ: token.T_Mul, Lit: "*"}
	_div          = token.Token{Typ: token.T_Div, Lit: "/"}
	_eof          = token.Token{Typ: token.T_Eof, Lit: "EOF"}
	_if           = token.Token{Typ: token.T_If, Lit: "IF"}
	_lcurly       = token.Token{Typ: token.T_LCurly, Lit: "{"}
	_rcurly       = token.Token{Typ: token.T_RCurly, Lit: "}"}
	_forever      = token.Token{Typ: token.T_Forever, Lit: "FOREVER"}
	_foreach      = token.Token{Typ: token.T_Foreach, Lit: "FOREACH"}
	_in           = token.Token{Typ: token.T_In, Lit: "IN"}
	_true         = token.Token{Typ: token.T_True, Lit: "TRUE"}
	_comma        = token.Token{Typ: token.T_Comma, Lit: ","}
	_match        = token.Token{Typ: token.T_Match, Lit: "MATCH"}
	_with         = token.Token{Typ: token.T_With, Lit: "WITH"}
	_single_arrow = token.Token{Typ: token.T_SingleArrow, Lit: "->"}
	_double_arrow = token.Token{Typ: token.T_DoubleArrow, Lit: "=>"}
	_dollar       = token.Token{Typ: token.T_Dollar, Lit: "$"}
	_lshift       = token.Token{Typ: token.T_BitLeftShift, Lit: "<<"}
	_neq          = token.Token{Typ: token.T_NotEq, Lit: "!="}
	_eqeq         = token.Token{Typ: token.T_DoubleEq, Lit: "=="}
	_geq          = token.Token{Typ: token.T_GtEq, Lit: ">="}
	_minusminus   = token.Token{Typ: token.T_MinusMinus, Lit: "--"}
	_plusplus     = token.Token{Typ: token.T_PlusPlus, Lit: "++"}
	_bitand       = token.Token{Typ: token.T_BitAnd, Lit: "&"}
	_bitor        = token.Token{Typ: token.T_BitOr, Lit: "|"}
	_and          = token.Token{Typ: token.T_And, Lit: "&&"}
	_or           = token.Token{Typ: token.T_Or, Lit: "||"}
	_bitxor       = token.Token{Typ: token.T_BitXor, Lit: "^"}
	_pluseq       = token.Token{Typ: token.T_PlusEq, Lit: "+="}
	_muleq        = token.Token{Typ: token.T_MulEq, Lit: "*="}
	_pub          = token.Token{Typ: token.T_Pub, Lit: "PUB"}
	_struct       = token.Token{Typ: token.T_Struct, Lit: "STRUCT"}
	_embeds       = token.Token{Typ: token.T_Embeds, Lit: "EMBEDS"}
	_dot          = token.Token{Typ: token.T_Dot, Lit: "."}
)

func TestScannerNext(t *testing.T) {
	input := "string name = \"Jennifer\""
	input += "\n"
	input += "const age uint = 44"
	input += "\n"
	input += "println(\"hello\")"
	input += "\n"
	input += "61 + 75 (12 *3) /  86 - 144"
	input += "if{}forever foreach _, item in "
	input += "true, 10.5 match with -> =>"
	input += "$ << != == >=- -- ++ & && | || ^ += *="
	input += "pub struct embeds ."

	s := New(input)
	got := []token.Token{}
	for {
		tok := s.Next()
		if tok.Typ == token.T_Eof {
			break
		}
		if tok.Typ != token.T_Whitespace {
			got = append(got, tok)
		}
	}
	want := []token.Token{
		_stringkw, _ident("name"), _eq, _string("Jennifer"), _const, _ident("age"), _uintKw, _eq,
		_number("44"), _println, _lparen, _string("hello"), _rparen, _number("61"), _plus,
		_number("75"), _lparen, _number("12"), _mul, _number("3"), _rparen,
		_div, _number("86"), _minus, _number("144"), _if, _lcurly, _rcurly,
		_forever, _foreach, _ident("_"), _comma, _ident("item"), _in,
		_true, _comma, _number("10.5"), _match, _with, _single_arrow,
		_double_arrow, _dollar, _lshift, _neq, _eqeq, _geq,
		_minus, _minusminus, _plusplus, _bitand, _and, _bitor,
		_or, _bitxor, _pluseq, _muleq, _pub, _struct, _embeds, _dot,
	}

	for i, v := range got {
		if i == len(want) {
			t.Errorf("out of bounds %d\n", i)
		}
		curWant := want[i]
		if curWant.Typ == token.T_Whitespace {
			continue
		}
		if v.Lit != curWant.Lit || v.Typ != curWant.Typ {
			t.Errorf("at index %d, wanted %v, but got %v\n", i, curWant, v)
		}
	}
}

/* 		_stringkw, _ws(" "), _ident("name"), _ws(" "), _ws(" "), _eq, _ws(" "), _string("Jennifer"),
_ws("\n"), _const, _ws(" "), _ident("age"), _ws(" "), _uintKw, _ws(" "), _eq, _ws(" "), _number("44"),
_ws("\n"), _println, _lparen, _string("hello"), _rparen, _ws("\n"), _number("61"), _ws(" "), _plus,
_ws(" "), _number("75"), _ws(" "), _lparen, _number("12"), _ws(" "), _mul, _number("3"), _rparen, _ws(" "),
_div, _ws("  "), _number("86"), _ws(" "), _minus, _ws(" "), _number("144"), _ws("\n"), _if, _lcurly, _rcurly,
_forever, _ws(" "), _foreach, _ws(" "), _ident("_"), _comma, _ws(" "), _ident("item"), _ws(" "), _in, _ws("\n"),
_true, _comma, _ws(" "), _number("10.5"), _ws(" "), _match, _ws(" "), _with, _ws(" "), _single_arrow, _ws(" "),
_double_arrow, _ws("\n"), _dollar, _ws(" "), _lshift, _ws(" "), _neq, _ws(" "), _eqeq, _ws(" "), _geq, _ws(" "),
_minus, _ws(" "), _minusminus, _ws(" "), _plusplus, _ws(" "), _bitand, _ws(" "), _and, _ws(" "), _bitor, _ws(" "),
_or, _ws(" "), _bitxor, _ws(" "), _pluseq, _ws(" "), _muleq,*/
