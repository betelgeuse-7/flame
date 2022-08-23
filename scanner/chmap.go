package scanner

import "flame/token"

var singleCharMap = map[byte]token.Token{
	'+':  {Typ: token.Plus, Lit: "+"},
	'-':  {Typ: token.Minus, Lit: "-"},
	'/':  {Typ: token.Div, Lit: "/"},
	'*':  {Typ: token.Mul, Lit: "*"},
	'!':  {Typ: token.Exclamation, Lit: "!"},
	'<':  {Typ: token.Lt, Lit: "<"},
	'>':  {Typ: token.Gt, Lit: ">"},
	'=':  {Typ: token.Eq, Lit: "="},
	'(':  {Typ: token.Lparen, Lit: "("},
	')':  {Typ: token.Rparen, Lit: ")"},
	'{':  {Typ: token.LCurly, Lit: "{"},
	'}':  {Typ: token.RCurly, Lit: "}"},
	',':  {Typ: token.Comma, Lit: ","},
	'.':  {Typ: token.Dot, Lit: "."},
	'$':  {Typ: token.Dollar, Lit: "$"},
	'\'': {Typ: token.SingleQuote, Lit: "'"},
	'%':  {Typ: token.Modulus, Lit: "%"},
	'#':  {Typ: token.Octothorp, Lit: "#"},
}

var doubleCharMap = map[string]token.Token{
	"--": {Typ: token.MinusMinus, Lit: "--"},
	"++": {Typ: token.PlusPlus, Lit: "++"},
	"+=": {Typ: token.PlusEq, Lit: "+="},
	"-=": {Typ: token.MinusEq, Lit: "-="},
	"*=": {Typ: token.MulEq, Lit: "*="},
	"/=": {Typ: token.DivEq, Lit: "/="},
	">=": {Typ: token.GtEq, Lit: ">="},
	"<=": {Typ: token.LtEq, Lit: "<="},
	"!=": {Typ: token.NotEq, Lit: "!="},
	"==": {Typ: token.DoubleEq, Lit: "=="},
	"&&": {Typ: token.And, Lit: "&&"},
	"||": {Typ: token.Or, Lit: "||"},
	"->": {Typ: token.SingleArrow, Lit: "->"},
	"=>": {Typ: token.DoubleArrow, Lit: "=>"},
}
