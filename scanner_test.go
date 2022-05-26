package main

import (
	"testing"
)

var (
	_var      = token{typ: t_Var, lit: "VAR"}
	_ws       = func(lit string) token { return token{typ: t_Whitespace, lit: lit} }
	_ident    = func(lit string) token { return token{typ: t_Ident, lit: lit} }
	_stringkw = token{typ: t_StringKw, lit: "STRINGKW"}
	_eq       = token{typ: t_Eq, lit: "="}
	_string   = func(lit string) token { return token{typ: t_String, lit: lit} }
	_const    = token{typ: t_Const, lit: "CONST"}
	_uintKw   = token{typ: t_UintKw, lit: "UINTKW"}
	_number   = func(lit string) token { return token{typ: t_Number, lit: lit} }
	_println  = token{typ: t_Println, lit: "PRINTLN"}
	_lparen   = token{typ: t_Lparen, lit: "("}
	_rparen   = token{typ: t_Rparen, lit: ")"}
	_plus     = token{typ: t_Plus, lit: "+"}
	_minus    = token{typ: t_Minus, lit: "-"}
	_mul      = token{typ: t_Mul, lit: "*"}
	_div      = token{typ: t_Div, lit: "/"}
	_eof      = token{typ: t_Eof, lit: "EOF"}
)

func TestScannerNext(t *testing.T) {
	input := "var name string = \"Jennifer\""
	input += "\n"
	input += "const age uint = 44"
	input += "\n"
	input += "println(\"hello\")"
	input += "\n"
	input += "61 + 75 (12 *3) /  86 - 144"

	s := newScanner(input)
	got := []token{}
	for {
		tok := s.next()
		if tok.typ == t_Eof {
			break
		}
		got = append(got, tok)
	}
	want := []token{
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
		if v.lit != curWant.lit || v.typ != curWant.typ {
			t.Errorf("at index %d, wanted %v, but got %v\n", i, curWant, v)
		}
	}
}
