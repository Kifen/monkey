package token

type TokenType string

type Position struct {
	Line   int
	Column int
}

type Token struct {
	Type     TokenType
	Literal  string
	Position Position
}

const (
	EOF       = "EOF"
	ILLEGAL   = "ILLEGAL"
	INT       = "INT"
	ASSIGN    = "ASSIGN"
	PLUS      = "PLUS"
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	LBRACE    = "LBRACE"
	RBRACE    = "RBRACE"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	IDENT     = "IDENT"
	BANG      = "BANG"
	ASTERISK  = "ASTERISK"
	SLASH     = "SLASH"
	LT        = "LT"
	GT        = "GT"
	MINUS     = "MINUS"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
	EQ        = "EQUAL"
	NOT_EQ    = "NOT-EQUAL"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
