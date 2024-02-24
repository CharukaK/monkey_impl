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
	IDENTIFIER = "IDENTIFIER" // identifiers
	INT        = "INTEGER"    // 12345

	// Operators
	Assign   = "="
	PLUS     = "+"
	BANG     = "!"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	// Comparative operators
	LT     = "<"
	GT     = ">"
	NOT_EQ = "!="
	EQ     = "=="

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
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "true"
	FALSE    = "false"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tokT, ok := keywords[ident]; ok {
		return tokT
	}

	return IDENTIFIER
}
