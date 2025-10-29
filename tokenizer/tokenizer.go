package tokenizer

import (
	"MonkeyInterpreter/token"
)

type Tokenizer struct {
	input        string
	position     int  // current position in input (current char)
	readPosition int  // character after current char (position)
	ch           byte // current char
}

// TODO: support for Unicode/UTF-8 characters
func NewTokenizer(input string) *Tokenizer {
	t := &Tokenizer{input: input}
	t.readChar()
	return t
}

func (t *Tokenizer) readChar() {
	if t.readPosition >= len(t.input) {
		t.ch = 0 // ASCII NUL character
	} else {
		t.ch = t.input[t.readPosition]
	}
	t.position = t.readPosition
	t.readPosition += 1
}

func (t *Tokenizer) NextToken() token.Token {
	var tok token.Token

	t.skipWhitespace()

	switch t.ch {
	case '=':
		if t.peekNextChar() == '=' {
			firstCh := t.ch
			t.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(firstCh) + string(t.ch)}
		} else {
			tok = newToken(token.ASSIGN, t.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, t.ch)
	case '(':
		tok = newToken(token.LPAREN, t.ch)
	case ')':
		tok = newToken(token.RPAREN, t.ch)
	case '{':
		tok = newToken(token.LBRACE, t.ch)
	case '}':
		tok = newToken(token.RBRACE, t.ch)
	case ',':
		tok = newToken(token.COMMA, t.ch)
	case '+':
		tok = newToken(token.PLUS, t.ch)
	case '!':
		if t.peekNextChar() == '=' {
			firstCh := t.ch
			t.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(firstCh) + string(t.ch)}
		} else {
			tok = newToken(token.BANG, t.ch)
		}
	case '-':
		tok = newToken(token.MINUS, t.ch)
	case '/':
		tok = newToken(token.SLASH, t.ch)
	case '*':
		tok = newToken(token.ASTERISK, t.ch)
	case '<':
		tok = newToken(token.LESS_THAN, t.ch)
	case '>':
		tok = newToken(token.GREATER_THAN, t.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(t.ch) {
			tok.Literal = t.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(t.ch) {
			tok.Type = token.INT
			tok.Literal = t.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, t.ch)
		}
	}
	t.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (t *Tokenizer) readIdentifier() string {
	pos := t.position
	for isLetter(t.ch) {
		t.readChar()
	}
	return t.input[pos:t.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (t *Tokenizer) skipWhitespace() {
	for t.ch == ' ' || t.ch == '\t' || t.ch == '\n' || t.ch == '\r' {
		t.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (t *Tokenizer) readNumber() string {
	pos := t.position
	for isDigit(t.ch) {
		t.readChar()
	}
	return t.input[pos:t.position]
}

func (t *Tokenizer) peekNextChar() byte {
	if t.readPosition >= len(t.input) {
		return 0
	}
	return t.input[t.readPosition]
}
