package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Tokens
const (
	ILLEGAL = "ILLEGAL" // tokens or characters that we can't figure out, let the parser know when to stop
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // identifiers
	INT   = "INT" // 12345

	// Operators
	Assign = "="
	PLUS   = "+"

	// Delimeters
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
