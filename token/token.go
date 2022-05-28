package token

type TokenType string

const (
	T_Eof        TokenType = "EOF"
	T_Illegal    TokenType = "ILLEGAL"
	T_Whitespace TokenType = "WHITESPACE"

	T_Ident  TokenType = "IDENT"
	T_String TokenType = "STRING"
	T_Number TokenType = "NUMBER"

	/* ****** KEYWORDS *********/
	T_Var      TokenType = "VAR"
	T_Const    TokenType = "CONST"
	T_StringKw TokenType = "STRINGKW"
	T_UintKw   TokenType = "UINTKW"
	T_IntKw    TokenType = "INTKW"
	T_BoolKw   TokenType = "BOOLKW"

	/* ****** OPERATORS *********/
	T_Lparen TokenType = "LPAREN"
	T_Rparen TokenType = "RPAREN"
	T_Plus   TokenType = "PLUS"
	T_Minus  TokenType = "MINUS"
	T_Mul    TokenType = "MUL"
	T_Div    TokenType = "DIV"
	T_Eq     TokenType = "EQ"

	/* ****** Built-ins *********/
	T_Println TokenType = "PRINTLN"
)

type Token struct {
	Typ TokenType
	Lit string
}
