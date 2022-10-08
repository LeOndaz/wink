package lexer

import (
	"lango/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
			let two = 2;
			let sixteen = 16;

			let add = fn (x, y) {
				x + y;
			}
			let result = add(two, sixteen);

			5 < 10 <= 20 / 2;
			let six = 6 != (5 * 2);
`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// let two = 2;
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENTIFIER, expectedLiteral: "two"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "2"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},

		// let sixteen = 16;
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENTIFIER, expectedLiteral: "sixteen"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "16"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},

		// let add = fn (x , y) { x + y }
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENTIFIER, expectedLiteral: "add"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.FUNCTION, expectedLiteral: "fn"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.IDENTIFIER, expectedLiteral: "x"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.IDENTIFIER, expectedLiteral: "y"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.LBRACE, expectedLiteral: "{"},
		{expectedType: token.IDENTIFIER, expectedLiteral: "x"},
		{expectedType: token.PLUS, expectedLiteral: "+"},
		{expectedType: token.IDENTIFIER, expectedLiteral: "y"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.RBRACE, expectedLiteral: "}"},

		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENTIFIER, expectedLiteral: "result"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.IDENTIFIER, expectedLiteral: "add"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.IDENTIFIER, expectedLiteral: "two"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.IDENTIFIER, expectedLiteral: "sixteen"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},

		{expectedType: token.INT, expectedLiteral: "5"},
		{expectedType: token.LT, expectedLiteral: "<"},
		{expectedType: token.INT, expectedLiteral: "10"},
		{expectedType: token.LT, expectedLiteral: "<"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "20"},
		{expectedType: token.FSLASH, expectedLiteral: "/"},
		{expectedType: token.INT, expectedLiteral: "2"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},

		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENTIFIER, expectedLiteral: "six"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "6"},
		{expectedType: token.EXCLAMATION, expectedLiteral: "!"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.INT, expectedLiteral: "5"},
		{expectedType: token.ASTERISK, expectedLiteral: "*"},
		{expectedType: token.INT, expectedLiteral: "2"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
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

		// wrong literal
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
