package parser

import "flame/token"

var dataTypeKws = []token.TokenType{
	token.T_VoidKw, token.T_StringKw, token.T_UintKw, token.T_Uint32Kw,
	token.T_IntKw, token.T_Int32Kw, token.T_Float32Kw, token.T_Float64Kw,
	token.T_BoolKw,
}
