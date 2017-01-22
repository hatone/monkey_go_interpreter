package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGGAL = "ILLEGAL"
	EOF      = "EOF"

	// identifirs + literals
	IDENT = "IDENT" // add, foobar, x,y
	INT   = "INT"

	// Operaters
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
