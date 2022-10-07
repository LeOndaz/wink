package lexer

import "lango/token"

type Lexer struct {
	input string

	position int
	// ch = input[pos]
	ch byte

	// readPosition = position + 1
	readPosition int
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) skipWhitespace() {
	// ignore all space characters
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readChar() {
	// read the next character from the lexer, applying all conditions

	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var t token.Token

	lexer.skipWhitespace()

	switch lexer.ch {
	case '+':
		t = newToken(token.PLUS, lexer.ch)
	case '=':
		t = newToken(token.ASSIGN, lexer.ch)
	case ';':
		t = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		t = newToken(token.LPAREN, lexer.ch)
	case ')':
		t = newToken(token.RPAREN, lexer.ch)
	case '{':
		t = newToken(token.LBRACE, lexer.ch)
	case '}':
		t = newToken(token.RBRACE, lexer.ch)
	case ',':
		t = newToken(token.COMMA, lexer.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(lexer.ch) {
			t.Literal = lexer.readIdentifier()
			t.Type = token.GetIdentifierToken(t.Literal)
			return t // needed because we've read the identifier already
		} else if isDigit(lexer.ch) {
			t.Type = token.INT
			t.Literal = lexer.readDigit()
			return t // early exit again
		} else {
			t = newToken(token.ILLEGAL, lexer.ch) // not a digit, not a string, wtf is it?
		}
	}

	lexer.readChar()
	return t
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) readIdentifier() string {
	oldPosition := lexer.position

	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[oldPosition:lexer.position]
}

func (lexer *Lexer) readDigit() string {
	oldPosition := lexer.position

	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[oldPosition:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '$'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
