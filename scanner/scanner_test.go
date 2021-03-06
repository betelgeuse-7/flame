package scanner

import (
	"flame/token"
	"testing"
)

var (
	_ident        = func(Lit string) token.Token { return token.Token{Typ: token.T_Ident, Lit: Lit} }
	_stringkw     = token.Token{Typ: token.T_StringKw, Lit: "string"}
	_octothorp    = token.Token{Typ: token.T_Octothorp, Lit: "#"}
	_eq           = token.Token{Typ: token.T_Eq, Lit: "="}
	_string       = func(Lit string) token.Token { return token.Token{Typ: token.T_String, Lit: Lit} }
	_i64          = func(lit string) token.Token { return token.Token{Typ: token.T_Int, Lit: lit} }
	_f64          = func(lit string) token.Token { return token.Token{Typ: token.T_Float64, Lit: lit} }
	_u64          = func(lit string) token.Token { return token.Token{Typ: token.T_Uint, Lit: lit} }
	_uintKw       = token.Token{Typ: token.T_UintKw, Lit: "uint"}
	_println      = token.Token{Typ: token.T_PrintlnFn, Lit: "printlnfn"}
	_lparen       = token.Token{Typ: token.T_Lparen, Lit: "("}
	_rparen       = token.Token{Typ: token.T_Rparen, Lit: ")"}
	_plus         = token.Token{Typ: token.T_Plus, Lit: "+"}
	_minus        = token.Token{Typ: token.T_Minus, Lit: "-"}
	_mul          = token.Token{Typ: token.T_Mul, Lit: "*"}
	_div          = token.Token{Typ: token.T_Div, Lit: "/"}
	_eof          = token.Token{Typ: token.T_Eof, Lit: "eof"}
	_if           = token.Token{Typ: token.T_If, Lit: "if"}
	_lcurly       = token.Token{Typ: token.T_LCurly, Lit: "{"}
	_rcurly       = token.Token{Typ: token.T_RCurly, Lit: "}"}
	_forever      = token.Token{Typ: token.T_Forever, Lit: "forever"}
	_foreach      = token.Token{Typ: token.T_Foreach, Lit: "foreach"}
	_in           = token.Token{Typ: token.T_In, Lit: "in"}
	_bool         = func(lit string) token.Token { return token.Token{Typ: token.T_Bool, Lit: lit} }
	_comma        = token.Token{Typ: token.T_Comma, Lit: ","}
	_match        = token.Token{Typ: token.T_Match, Lit: "match"}
	_with         = token.Token{Typ: token.T_With, Lit: "with"}
	_single_arrow = token.Token{Typ: token.T_SingleArrow, Lit: "->"}
	_double_arrow = token.Token{Typ: token.T_DoubleArrow, Lit: "=>"}
	_dollar       = token.Token{Typ: token.T_Dollar, Lit: "$"}
	_neq          = token.Token{Typ: token.T_NotEq, Lit: "!="}
	_eqeq         = token.Token{Typ: token.T_DoubleEq, Lit: "=="}
	_geq          = token.Token{Typ: token.T_GtEq, Lit: ">="}
	_minusminus   = token.Token{Typ: token.T_MinusMinus, Lit: "--"}
	_plusplus     = token.Token{Typ: token.T_PlusPlus, Lit: "++"}
	_and          = token.Token{Typ: token.T_And, Lit: "&&"}
	_or           = token.Token{Typ: token.T_Or, Lit: "||"}
	_pluseq       = token.Token{Typ: token.T_PlusEq, Lit: "+="}
	_muleq        = token.Token{Typ: token.T_MulEq, Lit: "*="}
	_pub          = token.Token{Typ: token.T_Pub, Lit: "pub"}
	_struct       = token.Token{Typ: token.T_Struct, Lit: "struct"}
	_embeds       = token.Token{Typ: token.T_Embeds, Lit: "embeds"}
	_dot          = token.Token{Typ: token.T_Dot, Lit: "."}
	_comment      = func(lit string) token.Token { return token.Token{Typ: token.T_Comment, Lit: lit} }
	_illegal      = func(lit string) token.Token { return token.Token{Typ: token.T_Illegal, Lit: lit} }
	_single_quote = token.Token{Typ: token.T_SingleQuote, Lit: "'"}
	_modulus      = token.Token{Typ: token.T_Modulus, Lit: "%"}
)

func TestScannerNext(t *testing.T) {
	input := "string name = \"Jennifer\""
	input += "\n"
	input += "#uint age = 44"
	input += "\n"
	input += "println(\"hello\")"
	input += "\n"
	input += "61 + 75 (12 *3) /  86 - 144"
	input += "if{}forever foreach _, item in "
	input += "true, 10.5 match with -> =>"
	input += "$ != == >=- -- ++ && || += *="
	input += "// this is a comment\n"
	input += "pub struct embeds ."
	input += "\"????\" #string ???? = \"blood\""
	input += "people'1 5 % 3 -178.6 -168"

	s := New(input)
	// we expect s.y to be 1 in the beginning
	expectSDotY := 1
	if s.y != expectSDotY {
		t.Errorf("expected s.y to be %d, but got %d\n", expectSDotY, s.y)
	}
	got := []token.Token{}
	for {
		tok := s.Next()
		if tok.Typ == token.T_Eof {
			break
		}
		got = append(got, tok)
	}
	want := []token.Token{
		_stringkw, _ident("name"), _eq, _string("Jennifer"), _octothorp, _uintKw, _ident("age"), _eq,
		_u64("44"), _println, _lparen, _string("hello"), _rparen, _u64("61"), _plus,
		_u64("75"), _lparen, _u64("12"), _mul, _u64("3"), _rparen,
		_div, _u64("86"), _minus, _u64("144"), _if, _lcurly, _rcurly,
		_forever, _foreach, _ident("_"), _comma, _ident("item"), _in,
		_bool("true"), _comma, _f64("10.5"), _match, _with, _single_arrow,
		_double_arrow, _dollar, _neq, _eqeq, _geq,
		_minus, _minusminus, _plusplus, _and,
		_or, _pluseq, _muleq, _comment(" this is a comment"), _pub, _struct, _embeds, _dot,
		_string("????"), _octothorp, _stringkw, _illegal("????"), _eq, _string("blood"),
		_ident("people"), _single_quote, _u64("1"), _u64("5"), _modulus, _u64("3"), _f64("-178.6"), _i64("-168"),
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
	// we expect s.y to be 5 in the end
	expectSDotY = 5
	if s.y != expectSDotY {
		t.Errorf("expected s.y to be %d, but got %d\n", expectSDotY, s.y)
	}
}
