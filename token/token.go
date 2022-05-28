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
	T_VoidKw    TokenType = "VOIDKW"
	T_StringKw  TokenType = "STRINGKW"
	T_UintKw    TokenType = "UINTKW"
	T_Uint32Kw  TokenType = "UINT32KW"
	T_IntKw     TokenType = "INTKW"
	T_Int32Kw   TokenType = "INT32KW"
	T_Float64Kw TokenType = "FLOAT64KW"
	T_Float32Kw TokenType = "FLOAT32KW"
	T_BoolKw    TokenType = "BOOLKW"

	T_Var      TokenType = "VAR"
	T_Const    TokenType = "CONST"
	T_Return   TokenType = "RETURN"
	T_Match    TokenType = "MATCH"
	T_With     TokenType = "WITH"
	T_If       TokenType = "IF"
	T_ElseIf   TokenType = "ELSEIF"
	T_Else     TokenType = "ELSE"
	T_True     TokenType = "TRUE"
	T_False    TokenType = "FALSE"
	T_Repeat   TokenType = "REPEAT"
	T_Forever  TokenType = "FOREVER"
	T_Foreach  TokenType = "FOREACH"
	T_While    TokenType = "WHILE"
	T_Break    TokenType = "BREAK"
	T_Continue TokenType = "CONTINUE"
	T_Struct   TokenType = "STRUCT"
	T_Embeds   TokenType = "EMBEDS"
	T_Pub      TokenType = "PUB"
	T_Pkg      TokenType = "PKG"
	T_Import   TokenType = "IMPORT"

	/* ****** OPERATORS *********/
	T_Lparen        TokenType = "LPAREN"
	T_Rparen        TokenType = "RPAREN"
	T_LSquareParen  TokenType = "LSQUAREPAREN"
	T_RSquareParen  TokenType = "RSQUAREPAREN"
	T_LCurly        TokenType = "LCURLY"
	T_RCurly        TokenType = "RCURLY"
	T_Plus          TokenType = "PLUS"
	T_PlusPlus      TokenType = "PLUSPLUS"
	T_Minus         TokenType = "MINUS"
	T_MinusMinus    TokenType = "MINUSMINUS"
	T_Mul           TokenType = "MUL"
	T_Div           TokenType = "DIV"
	T_Eq            TokenType = "EQ"
	T_PlusEq        TokenType = "PLUSEQ"
	T_MinusEq       TokenType = "MINUSEQ"
	T_MulEq         TokenType = "MULEQ"
	T_DivEq         TokenType = "DIVEQ"
	T_NotEq         TokenType = "NOTEQ"
	T_DoubleEq      TokenType = "DOUBLEEQ"
	T_BitLeftShift  TokenType = "BITLEFTSHIFT"  // <<
	T_BitRightShift TokenType = "BITRIGHTSHIFT" // >>
	T_BitAnd        TokenType = "BITAND"        // &
	T_BitOr         TokenType = "BITOR"         // |
	T_BitXor        TokenType = "BITXOR"        // ^
	T_And           TokenType = "AND"
	T_Or            TokenType = "OR"
	T_Dot           TokenType = "DOT"
	T_Gt            TokenType = "GT"
	T_Lt            TokenType = "LT"
	T_Dollar        TokenType = "DOLLAR"
	T_Ampersand     TokenType = "AMPERSAND"
	T_Colon         TokenType = "COLON"
	T_Comma         TokenType = "COMMA"
	T_Exclamation   TokenType = "EXCLAMATION"
	T_SingleArrow   TokenType = "SINGLEARROW" // ->
	T_DoubleArrow   TokenType = "DOUBLEARROW" // =>

	/* ****** Built-ins *********/
	T_PrintlnFn TokenType = "PRINTLNFN"
	T_SetFn     TokenType = "SETFN"
	T_AppendFn  TokenType = "APPENDFN"
	T_DeleteFn  TokenType = "DELETEFN"
	T_PopFn     TokenType = "POPFN"
)

type Token struct {
	Typ TokenType
	Lit string
}
