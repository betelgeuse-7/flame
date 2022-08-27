package scanner

import "flame/token"

var keywordMap = map[string]token.Token{
	"string":   {Typ: token.StringKw},
	"char":     {Typ: token.CharKw},
	"uint":     {Typ: token.UintKw},
	"int":      {Typ: token.IntKw},
	"float":    {Typ: token.FloatKw},
	"bool":     {Typ: token.BoolKw},
	"return":   {Typ: token.Return},
	"if":       {Typ: token.If},
	"else":     {Typ: token.Else},
	"elseif":   {Typ: token.ElseIf},
	"true":     {Typ: token.Bool},
	"false":    {Typ: token.Bool},
	"repeat":   {Typ: token.Repeat},
	"forever":  {Typ: token.Forever},
	"while":    {Typ: token.While},
	"break":    {Typ: token.Break},
	"continue": {Typ: token.Continue},
	"println":  {Typ: token.PrintlnFn},
	"strLen":   {Typ: token.StrLenFn},
	"push":     {Typ: token.PushFn},
}
