package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string // literal文字
}


// literal identifier
const(
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "ASSIGN"
	PLUS = "+"

	COMMA = "."
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
)


