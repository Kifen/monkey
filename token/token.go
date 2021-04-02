package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	EOF = "EOF"
	ILLEGAL = "ILLEGAL"
	INT = "INT"
	ASSIGN = "="
	PLUS = "+"
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	FUNCTION = "FUNCTION"
	LET = "LET"
	IDENT = "IDENT"
)

var keywords = map[string] TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}