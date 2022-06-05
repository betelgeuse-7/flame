package scanner

import "flame/token"

var singleCharMap = map[byte]token.Token{
	'+':  {Typ: token.T_Plus, Lit: "+"},
	'-':  {Typ: token.T_Minus, Lit: "-"},
	'/':  {Typ: token.T_Div, Lit: "/"},
	'*':  {Typ: token.T_Mul, Lit: "*"},
	'!':  {Typ: token.T_Exclamation, Lit: "!"},
	'&':  {Typ: token.T_Ampersand, Lit: "&"},
	'|':  {Typ: token.T_BitOr, Lit: "|"},
	'<':  {Typ: token.T_Lt, Lit: "<"},
	'>':  {Typ: token.T_Gt, Lit: ">"},
	'=':  {Typ: token.T_Eq, Lit: "="},
	'^':  {Typ: token.T_BitXor, Lit: "^"},
	'(':  {Typ: token.T_Lparen, Lit: "("},
	')':  {Typ: token.T_Rparen, Lit: ")"},
	'{':  {Typ: token.T_LCurly, Lit: "{"},
	'}':  {Typ: token.T_RCurly, Lit: "}"},
	',':  {Typ: token.T_Comma, Lit: ","},
	'.':  {Typ: token.T_Dot, Lit: "."},
	'$':  {Typ: token.T_Dollar, Lit: "$"},
	'\'': {Typ: token.T_SingleQuote, Lit: "'"},
	'%':  {Typ: token.T_Modulus, Lit: "%"},
}

var doubleCharMap = map[string]token.Token{
	"--": {Typ: token.T_MinusMinus, Lit: "--"},
	"++": {Typ: token.T_PlusPlus, Lit: "++"},
	"+=": {Typ: token.T_PlusEq, Lit: "+="},
	"-=": {Typ: token.T_MinusEq, Lit: "-="},
	"*=": {Typ: token.T_MulEq, Lit: "*="},
	"/=": {Typ: token.T_DivEq, Lit: "/="},
	">=": {Typ: token.T_GtEq, Lit: ">="},
	"<=": {Typ: token.T_LtEq, Lit: "<="},
	"!=": {Typ: token.T_NotEq, Lit: "!="},
	"==": {Typ: token.T_DoubleEq, Lit: "=="},
	"<<": {Typ: token.T_BitLeftShift, Lit: "<<"},
	">>": {Typ: token.T_BitRightShift, Lit: ">>"},
	"&&": {Typ: token.T_And, Lit: "&&"},
	"||": {Typ: token.T_Or, Lit: "||"},
	"->": {Typ: token.T_SingleArrow, Lit: "->"},
	"=>": {Typ: token.T_DoubleArrow, Lit: "=>"},
}
