package lexer

import (
	"testing"
	"interpreter/token"
)

// TestNextToken: 
func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	// A slice of anonoymous structs with fields expectedType
	//  and expectedLiteral.
	// Initialzes the tests slice with the test cases below. based on input.
	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	// Iterate as long as there are tests in the tests slice
	// initializes i and tt, range tests expression iterates over the tests slice.
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got %q", i, tt.expectedType, tok.Type)
		} else if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		} 
	}
}