package token

type TokenType string

const (
	Eof     TokenType = "EOF"
	Illegal TokenType = "ILLEGAL"
	Newline TokenType = "NEWLINE"

	Ident   TokenType = "IDENT"
	String  TokenType = "STRING"
	Char    TokenType = "CHAR"
	Uint    TokenType = "UINT"
	Int     TokenType = "INT"
	Float   TokenType = "FLOAT"
	Bool    TokenType = "BOOL"
	Comment TokenType = "COMMENT"

	/* ****** KEYWORDS *********/
	StringKw TokenType = "string"
	CharKw   TokenType = "char"
	UintKw   TokenType = "uint"
	IntKw    TokenType = "int"
	FloatKw  TokenType = "float"
	BoolKw   TokenType = "bool"

	Return   TokenType = "RETURN"
	If       TokenType = "IF"
	ElseIf   TokenType = "ELSEIF"
	Else     TokenType = "ELSE"
	True     TokenType = "TRUE"
	False    TokenType = "FALSE"
	Repeat   TokenType = "REPEAT"
	Forever  TokenType = "FOREVER"
	While    TokenType = "WHILE"
	Break    TokenType = "BREAK"
	Continue TokenType = "CONTINUE"
	Phunc    TokenType = "PHUNC"

	/* ****** OPERATORS *********/
	Lparen       TokenType = "LPAREN"
	Rparen       TokenType = "RPAREN"
	LSquareParen TokenType = "LSQUAREPAREN"
	RSquareParen TokenType = "RSQUAREPAREN"
	LCurly       TokenType = "LCURLY"
	RCurly       TokenType = "RCURLY"
	Plus         TokenType = "PLUS"
	PlusPlus     TokenType = "PLUSPLUS"
	Minus        TokenType = "MINUS"
	MinusMinus   TokenType = "MINUSMINUS"
	Mul          TokenType = "MUL"
	Div          TokenType = "DIV"
	Modulus      TokenType = "MODULUS"
	Eq           TokenType = "EQ"
	PlusEq       TokenType = "PLUSEQ"
	MinusEq      TokenType = "MINUSEQ"
	MulEq        TokenType = "MULEQ"
	DivEq        TokenType = "DIVEQ"
	NotEq        TokenType = "NOTEQ"
	DoubleEq     TokenType = "DOUBLEEQ"
	And          TokenType = "AND"
	Or           TokenType = "OR"
	Dot          TokenType = "DOT"
	Gt           TokenType = "GT"
	Lt           TokenType = "LT"
	GtEq         TokenType = "GTEQ" // >=
	LtEq         TokenType = "LTEQ" // <=
	Dollar       TokenType = "DOLLAR"
	Colon        TokenType = "COLON"
	SingleQuote  TokenType = "SINGLE_QUOTE"
	Comma        TokenType = "COMMA"
	Exclamation  TokenType = "EXCLAMATION"
	SingleArrow  TokenType = "SINGLEARROW" // ->
	DoubleArrow  TokenType = "DOUBLEARROW" // =>
	Octothorp    TokenType = "OCTOTHORP"   // #

	/* ****** Built-ins *********/
	PrintlnFn TokenType = "PRINTLNFN"
	PushFn    TokenType = "PUSHFN"
	DeleteFn  TokenType = "DELETEFN"
	StrLenFn  TokenType = "STRLENFN"
)

type Token struct {
	Typ TokenType
	Lit string
	Pos Pos
}

// X -> column
// Y -> line
type Pos struct {
	X, Y int
}
