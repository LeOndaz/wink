package lexer

import "lango/token"

type Lexer struct {
	input string

	position int
	// ch = input[pos]
	ch byte

	readPosition int
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readChar() {
	lexer.skipWhitespace()

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

	switch lexer.ch {

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
	}

	lexer.readChar()
	return t
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
