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

	var currentChar = string(lexer.ch)

	switch lexer.ch {
	case '+':
		t = newToken(token.PLUS, currentChar)
	case '-':
		t = newToken(token.MINUS, currentChar)
	case '=':
		if lexer.peakChar() == '=' {
			lexer.readChar()
			t = newToken(token.EQ, currentChar+string(lexer.ch))
		} else {
			t = newToken(token.ASSIGN, currentChar)
		}
	case ';':
		t = newToken(token.SEMICOLON, currentChar)
	case '(':
		t = newToken(token.LPAREN, currentChar)
	case ')':
		t = newToken(token.RPAREN, currentChar)
	case '{':
		t = newToken(token.LBRACE, currentChar)
	case '}':
		t = newToken(token.RBRACE, currentChar)
	case ',':
		t = newToken(token.COMMA, currentChar)
	case '*':
		t = newToken(token.ASTERISK, currentChar)
	case '!':
		if lexer.peakChar() == '=' {
			lexer.readChar()
			t = newToken(token.NE, currentChar+string(lexer.ch))
		} else {
			t = newToken(token.EXCLAMATION, currentChar)
		}
	case '<':
		if lexer.peakChar() == '=' {
			lexer.readChar()
			t = newToken(token.LTE, currentChar+string(lexer.ch))
		} else {
			return newToken(token.LT, currentChar)
		}
	case '>':
		if lexer.peakChar() == '=' {
			lexer.readChar()
			t = newToken(token.GTE, currentChar+string(lexer.ch))
		} else {
			return newToken(token.GT, currentChar)
		}
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
			t = newToken(token.ILLEGAL, currentChar) // not a digit, not a string, wtf is it?
		}
	}

	lexer.readChar()
	return t
}

func newToken(tokenType token.Type, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
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

func (lexer *Lexer) peakChar() uint8 {
	return lexer.input[lexer.readPosition]
}
