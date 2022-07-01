package parser

import (
	"flame/token"
)

func isDataTypeKw(typ token.TokenType) bool {
	for _, v := range dataTypeKws {
		if typ == v {
			return true
		}
	}
	return false
}
