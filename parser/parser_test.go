package parser

import (
	"github.com/MoonShining/condition/lexer"
	"testing"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestIdentifierExpression(t *testing.T) {
	input := `true && (a == null) && (b in [1,2] || (c > 5)) || (d <= 6 || (f >= 4)) && g in [1,"2"]`
	l := lexer.New(input)
	p := New(l)
	stmt := p.Parse()
	checkParserErrors(t, p)
	t.Log(stmt)
}

func TestParseError(t *testing.T) {
	input := `a &a`
	l := lexer.New(input)
	p := New(l)
	p.Parse()
	if len(p.Errors()) == 0 {
		t.Fail()
	}
}
