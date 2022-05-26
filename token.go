package main

type tokenType string

const (
	t_Eof        tokenType = "EOF"
	t_Illegal    tokenType = "ILLEGAL"
	t_Whitespace tokenType = "WHITESPACE"

	t_Ident  tokenType = "IDENT"
	t_String tokenType = "STRING"
	t_Number tokenType = "NUMBER"

	/* ****** KEYWORDS *********/
	t_Var      tokenType = "VAR"
	t_Const    tokenType = "CONST"
	t_StringKw tokenType = "STRINGKW"
	t_UintKw   tokenType = "UINTKW"
	t_IntKw    tokenType = "INTKW"
	t_BoolKw   tokenType = "BOOLKW"

	/* ****** OPERATORS *********/
	t_Lparen tokenType = "LPAREN"
	t_Rparen tokenType = "RPAREN"
	t_Plus   tokenType = "PLUS"
	t_Minus  tokenType = "MINUS"
	t_Mul    tokenType = "MUL"
	t_Div    tokenType = "DIV"
	t_Eq     tokenType = "EQ"

	/* ****** Built-ins *********/
	t_Println tokenType = "PRINTLN"
)

type token struct {
	typ tokenType
	lit string
}
