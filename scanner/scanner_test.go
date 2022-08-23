package scanner

import (
	"flame/token"
	"testing"
)

var (
	_ident        = func(Lit string) token.Token { return token.Token{Typ: token.Ident, Lit: Lit} }
	_stringkw     = token.Token{Typ: token.StringKw, Lit: "string"}
	_octothorp    = token.Token{Typ: token.Octothorp, Lit: "#"}
	_eq           = token.Token{Typ: token.Eq, Lit: "="}
	_string       = func(Lit string) token.Token { return token.Token{Typ: token.String, Lit: Lit} }
	_i64          = func(lit string) token.Token { return token.Token{Typ: token.Int, Lit: lit} }
	_f64          = func(lit string) token.Token { return token.Token{Typ: token.Float, Lit: lit} }
	_u64          = func(lit string) token.Token { return token.Token{Typ: token.Uint, Lit: lit} }
	_uintKw       = token.Token{Typ: token.UintKw, Lit: "uint"}
	_println      = token.Token{Typ: token.PrintlnFn, Lit: "printlnfn"}
	_lparen       = token.Token{Typ: token.Lparen, Lit: "("}
	_rparen       = token.Token{Typ: token.Rparen, Lit: ")"}
	_plus         = token.Token{Typ: token.Plus, Lit: "+"}
	_minus        = token.Token{Typ: token.Minus, Lit: "-"}
	_mul          = token.Token{Typ: token.Mul, Lit: "*"}
	_div          = token.Token{Typ: token.Div, Lit: "/"}
	_eof          = token.Token{Typ: token.Eof, Lit: "eof"}
	_if           = token.Token{Typ: token.If, Lit: "if"}
	_lcurly       = token.Token{Typ: token.LCurly, Lit: "{"}
	_rcurly       = token.Token{Typ: token.RCurly, Lit: "}"}
	_forever      = token.Token{Typ: token.Forever, Lit: "forever"}
	_bool         = func(lit string) token.Token { return token.Token{Typ: token.Bool, Lit: lit} }
	_comma        = token.Token{Typ: token.Comma, Lit: ","}
	_single_arrow = token.Token{Typ: token.SingleArrow, Lit: "->"}
	_double_arrow = token.Token{Typ: token.DoubleArrow, Lit: "=>"}
	_dollar       = token.Token{Typ: token.Dollar, Lit: "$"}
	_neq          = token.Token{Typ: token.NotEq, Lit: "!="}
	_eqeq         = token.Token{Typ: token.DoubleEq, Lit: "=="}
	_geq          = token.Token{Typ: token.GtEq, Lit: ">="}
	_minusminus   = token.Token{Typ: token.MinusMinus, Lit: "--"}
	_plusplus     = token.Token{Typ: token.PlusPlus, Lit: "++"}
	_and          = token.Token{Typ: token.And, Lit: "&&"}
	_or           = token.Token{Typ: token.Or, Lit: "||"}
	_pluseq       = token.Token{Typ: token.PlusEq, Lit: "+="}
	_muleq        = token.Token{Typ: token.MulEq, Lit: "*="}
	_dot          = token.Token{Typ: token.Dot, Lit: "."}
	_comment      = func(lit string) token.Token { return token.Token{Typ: token.Comment, Lit: lit} }
	_illegal      = func(lit string) token.Token { return token.Token{Typ: token.Illegal, Lit: lit} }
	_single_quote = token.Token{Typ: token.SingleQuote, Lit: "'"}
	_modulus      = token.Token{Typ: token.Modulus, Lit: "%"}
	_newline      = token.Token{Typ: token.Newline, Lit: "\\n"}
)

func TestScannerNext(t *testing.T) {
	input := "string name = \"Jennifer\""
	input += "\n"
	input += "#uint age = 44"
	input += "\n"
	input += "println(\"hello\")"
	input += "\n"
	input += "61 + 75 (12 *3) /  86 - 144"
	input += "if{}forever _, item "
	input += "true, 10.5 -> =>"
	input += "$ != == >=- -- ++ && || += *="
	input += "// this is a comment\n"
	input += ". "
	input += "\"🙂\" #string 🩸 = \"blood\""
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
		if tok.Typ == token.Eof {
			break
		}
		got = append(got, tok)
	}
	want := []token.Token{
		_stringkw, _ident("name"), _eq, _string("Jennifer"), _newline,
		_octothorp, _uintKw, _ident("age"), _eq, _u64("44"), _newline,
		_println, _lparen, _string("hello"), _rparen, _newline, _u64("61"),
		_plus, _u64("75"), _lparen, _u64("12"), _mul, _u64("3"), _rparen,
		_div, _u64("86"), _minus, _u64("144"), _if, _lcurly, _rcurly,
		_forever, _ident("_"), _comma, _ident("item"), _bool("true"),
		_comma, _f64("10.5"), _single_arrow, _double_arrow, _dollar,
		_neq, _eqeq, _geq, _minus, _minusminus, _plusplus, _and, _or,
		_pluseq, _muleq, _comment(" this is a comment"), _newline, _dot, _string("🙂"),
		_octothorp, _stringkw, _illegal("🩸"), _eq, _string("blood"),
		_ident("people"), _single_quote, _u64("1"), _u64("5"), _modulus,
		_u64("3"), _f64("-178.6"), _i64("-168"),
	}
	for i, v := range got {
		if i == len(want) {
			t.Errorf("out of bounds %d\n", i)
			break
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

func TestNewline(t *testing.T) {
	input := "hello\nhe"
	want := []token.Token{
		{Typ: token.Ident},
		{Typ: token.Newline},
		{Typ: token.Ident},
		{Typ: token.Eof},
	}
	l := New(input)
	i := 0
	for {
		if i == len(want) {
			break
		}
		tok := l.Next()
		if tok.Typ != want[i].Typ {
			t.Errorf("error (index#%d): wanted: %+v, got: %+v\n", i, want[i], tok)
		}
		i++
	}
}
