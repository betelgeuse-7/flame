package parser

import "flame/token"

var dataTypeKws = []token.TokenType{
	token.StringKw, token.UintKw,
	token.IntKw, token.FloatKw,
	token.BoolKw,
}
