package lexer

import (
	"github.com/MoonShining/condition/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
(a && !(c || d))
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.AND, "&&"},
		{token.BANG, "!"},
		{token.LPAREN, "("},
		{token.IDENT, "c"},
		{token.OR, "||"},
		{token.IDENT, "d"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
