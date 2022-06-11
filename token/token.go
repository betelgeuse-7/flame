package token

type TokenType string

const (
	T_Eof     TokenType = "EOF"
	T_Illegal TokenType = "ILLEGAL"
	//T_Whitespace TokenType = "WHITESPACE"

	T_Ident   TokenType = "IDENT"
	T_String  TokenType = "STRING"
	T_Number  TokenType = "NUMBER"
	T_Comment TokenType = "COMMENT"

	/* ****** KEYWORDS *********/
	T_VoidKw    TokenType = "void"
	T_StringKw  TokenType = "string"
	T_UintKw    TokenType = "uint"
	T_Uint32Kw  TokenType = "u32"
	T_IntKw     TokenType = "int"
	T_Int32Kw   TokenType = "i32"
	T_Float64Kw TokenType = "float"
	T_Float32Kw TokenType = "f32"
	T_BoolKw    TokenType = "bool"

	T_Return   TokenType = "RETURN"
	T_Match    TokenType = "MATCH"
	T_With     TokenType = "WITH"
	T_If       TokenType = "IF"
	T_ElseIf   TokenType = "ELSEIF"
	T_Else     TokenType = "ELSE"
	T_True     TokenType = "true"
	T_False    TokenType = "false"
	T_Repeat   TokenType = "REPEAT"
	T_Forever  TokenType = "FOREVER"
	T_Foreach  TokenType = "FOREACH"
	T_In       TokenType = "IN"
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
	T_Modulus       TokenType = "MODULUS"
	T_Eq            TokenType = "EQ"
	T_PlusEq        TokenType = "PLUSEQ"
	T_MinusEq       TokenType = "MINUSEQ"
	T_MulEq         TokenType = "MULEQ"
	T_DivEq         TokenType = "DIVEQ"
	T_NotEq         TokenType = "NOTEQ"
	T_DoubleEq      TokenType = "DOUBLEEQ"
	T_BitLeftShift  TokenType = "BITLEFTSHIFT"  // <<
	T_BitRightShift TokenType = "BITRIGHTSHIFT" // >>
	T_Ampersand     TokenType = "AMPERSAND"     // &
	T_BitOr         TokenType = "BITOR"         // |
	T_BitXor        TokenType = "BITXOR"        // ^
	T_And           TokenType = "AND"
	T_Or            TokenType = "OR"
	T_Dot           TokenType = "DOT"
	T_Gt            TokenType = "GT"
	T_Lt            TokenType = "LT"
	T_GtEq          TokenType = "GTEQ" // >=
	T_LtEq          TokenType = "LTEQ" // <=
	T_Dollar        TokenType = "DOLLAR"
	T_Colon         TokenType = "COLON"
	T_SingleQuote   TokenType = "SINGLE_QUOTE"
	T_Comma         TokenType = "COMMA"
	T_Exclamation   TokenType = "EXCLAMATION"
	T_SingleArrow   TokenType = "SINGLEARROW" // ->
	T_DoubleArrow   TokenType = "DOUBLEARROW" // =>
	T_Octothorp     TokenType = "OCTOTHORP"   // #

	/* ****** Built-ins *********/
	T_PrintlnFn TokenType = "PRINTLNFN"
	T_AppendFn  TokenType = "APPENDFN"
	T_DeleteFn  TokenType = "DELETEFN"
	T_PopFn     TokenType = "POPFN"
	T_StrLenFn  TokenType = "STRLENFN"
)

type Token struct {
	Typ TokenType
	Lit string
	Pos TokenPos
}

// X -> column
// Y -> line
type TokenPos struct {
	X, Y int
}
