package lexer

import (
	"bufio"
	"io"
	"unicode"

	"github.com/Kifen/monkey/token"
)

type Lexer struct {
	r      rune
	pos    token.Position
	reader *bufio.Reader
}

func New(r io.Reader) *Lexer {
	return &Lexer{
		pos:    token.Position{Line: 1, Column: 0},
		reader: bufio.NewReader(r),
	}
}

func (l *Lexer) resetPosition() {
	l.pos.Line++
	l.pos.Column = 0
}

func newToken(tokenType token.TokenType, ch rune, pos token.Position) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Position: pos}
}

func (l *Lexer) makeTwoCharToken(r rune) token.Token {
	var tok token.Token
	char := l.peekChar()

	switch r {
	case '!':
		if char == '=' {
			tok = token.Token{Type: token.NOT_EQ, Literal: string(r) + string(char)}
		} else {
			l.backup()
			tok = newToken(token.BANG, r, l.pos)
		}
	case '=':
		if char == '=' {
			tok = token.Token{Type: token.EQ, Literal: string(r) + string(char)}
		} else {
			l.backup()
			tok = newToken(token.ASSIGN, r, l.pos)
		}
	}

	return tok
}

func (l *Lexer) peekChar() rune {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		panic(err)
	}

	return r
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				tok.Literal = ""
				tok.Type = token.EOF
				return tok
			}

			// at this point there isn't much we can do, and the compiler
			// should just return the raw error to the user
			panic(err)
		}

		// update the column to the position of the newly read in rune
		l.pos.Column++

		switch r {
		case '\n':
			l.resetPosition()
		case '=':
			return l.makeTwoCharToken(r)
		case '!':
			return l.makeTwoCharToken(r)
		case ';':
			return newToken(token.SEMICOLON, r, l.pos)
		case '(':
			return newToken(token.LPAREN, r, l.pos)
		case ')':
			return newToken(token.RPAREN, r, l.pos)
		case ',':
			return newToken(token.COMMA, r, l.pos)
		case '+':
			return newToken(token.PLUS, r, l.pos)
		case '{':
			return newToken(token.LBRACE, r, l.pos)
		case '}':
			return newToken(token.RBRACE, r, l.pos)
		case '-':
			return newToken(token.MINUS, r, l.pos)
		case '/':
			return newToken(token.SLASH, r, l.pos)
		case '*':
			return newToken(token.ASTERISK, r, l.pos)
		case '<':
			return newToken(token.LT, r, l.pos)
		case '>':
			return newToken(token.GT, r, l.pos)
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				tok.Literal = l.readInt()
				tok.Type = token.INT
				tok.Position = startPos
				return tok
			} else if unicode.IsLetter(r) {
				startPos := l.pos
				l.backup()
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				tok.Position = startPos
				return tok
			} else {
				return newToken(token.ILLEGAL, r, token.Position{})
			}
		}
	}
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.Column--
}

func (l *Lexer) readInt() string {
	var lit string

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the int
				return lit
			}
		}

		l.pos.Column++
		if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) readIdentifier() string {
	var lit string

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the int
				return lit
			}
		}

		l.pos.Column++
		if unicode.IsLetter(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}
