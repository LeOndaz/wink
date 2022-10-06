package lexer

import (
	"lango/token"
	"testing"
)

func testNextToken(t *testing.T) {
	var input string = "+!@+=="

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{expectedType: token.PLUS, expectedLiteral: "+"},
		{expectedType: token.ILLEGAL, expectedLiteral: "!"},
		{expectedType: token.ILLEGAL, expectedLiteral: "@"},
		{expectedType: token.PLUS, expectedLiteral: "+"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
	}

	lexer := New(input)

	for i, tc := range tests {
		currentToken := lexer.NextToken()

		// wrong type
		if currentToken.Type != tc.expectedType {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q, got=%q",
				i,
				tc.expectedType,
				currentToken.Type,
			)
		}

		if currentToken.Literal != tc.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong. expected=%q, got=%q",
				i,
				tc.expectedLiteral,
				currentToken.Literal,
			)
		}

	}
}
