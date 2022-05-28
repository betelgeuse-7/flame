package scanner

import (
	"flame/token"
	"testing"
)

var (
	_var      = token.Token{Typ: token.T_Var, Lit: "VAR"}
	_ws       = func(Lit string) token.Token { return token.Token{Typ: token.T_Whitespace, Lit: Lit} }
	_ident    = func(Lit string) token.Token { return token.Token{Typ: token.T_Ident, Lit: Lit} }
	_stringkw = token.Token{Typ: token.T_StringKw, Lit: "STRINGKW"}
	_eq       = token.Token{Typ: token.T_Eq, Lit: "="}
	_string   = func(Lit string) token.Token { return token.Token{Typ: token.T_String, Lit: Lit} }
	_const    = token.Token{Typ: token.T_Const, Lit: "CONST"}
	_uintKw   = token.Token{Typ: token.T_UintKw, Lit: "UINTKW"}
	_number   = func(Lit string) token.Token { return token.Token{Typ: token.T_Number, Lit: Lit} }
	_println  = token.Token{Typ: token.T_Println, Lit: "PRINTLN"}
	_lparen   = token.Token{Typ: token.T_Lparen, Lit: "("}
	_rparen   = token.Token{Typ: token.T_Rparen, Lit: ")"}
	_plus     = token.Token{Typ: token.T_Plus, Lit: "+"}
	_minus    = token.Token{Typ: token.T_Minus, Lit: "-"}
	_mul      = token.Token{Typ: token.T_Mul, Lit: "*"}
	_div      = token.Token{Typ: token.T_Div, Lit: "/"}
	_eof      = token.Token{Typ: token.T_Eof, Lit: "EOF"}
)

func TestScannerNext(t *testing.T) {
	input := "var name string = \"Jennifer\""
	input += "\n"
	input += "const age uint = 44"
	input += "\n"
	input += "println(\"hello\")"
	input += "\n"
	input += "61 + 75 (12 *3) /  86 - 144"

	s := New(input)
	got := []token.Token{}
	for {
		tok := s.Next()
		if tok.Typ == token.T_Eof {
			break
		}
		got = append(got, tok)
	}
	want := []token.Token{
		_var, _ws(" "), _ident("name"), _ws(" "), _stringkw, _ws(" "), _eq, _ws(" "), _string("Jennifer"),
		_ws("\n"), _const, _ws(" "), _ident("age"), _ws(" "), _uintKw, _ws(" "), _eq, _ws(" "), _number("44"),
		_ws("\n"), _println, _lparen, _string("hello"), _rparen, _ws("\n"), _number("61"), _ws(" "), _plus,
		_ws(" "), _number("75"), _ws(" "), _lparen, _number("12"), _ws(" "), _mul, _number("3"), _rparen, _ws(" "),
		_div, _ws("  "), _number("86"), _ws(" "), _minus, _ws(" "), _number("144"), _ws(""),
	}

	for i, v := range got {
		if i == len(want) {
			t.Errorf("out of bounds %d\n", i)
		}
		curWant := want[i]
		if v.Lit != curWant.Lit || v.Typ != curWant.Typ {
			t.Errorf("at index %d, wanted %v, but got %v\n", i, curWant, v)
		}
	}
}
