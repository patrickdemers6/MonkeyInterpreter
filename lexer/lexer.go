package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()

			tok.Type = token.EQ
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()

			tok.Type = token.NOT_EQ
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()

			tok.Type = token.GTE
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()

			tok.Type = token.LTE
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// if not one of the tokens above and we have a letter, read the identifier
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	start_position := l.position
	// read letters while in current identifier
	for isLetter(l.ch) {
		l.readChar()
	}

	// return from starting position to end position of identifier
	return l.input[start_position:l.position]
}

func (l *Lexer) readNumber() string {
	start_position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[start_position:l.position]
}

func isLetter(token byte) bool {
	// in code, only a-z, A-Z, and _ are considered letters
	return (token >= 'a' && token <= 'z') || (token >= 'A' && token <= 'Z') || token == '_'
}

func isDigit(token byte) bool {
	return (token >= '0' && token <= '9')
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
