package run

import (
	"github.com/MoonShining/condition/lexer"
	"github.com/MoonShining/condition/parser"
	"testing"
)

func TestVisitor(t *testing.T) {
	input := `(false || a==1) && (b==2||c==3) && (d in ["hello"])`
	l := lexer.New(input)
	p := parser.New(l)
	exp := p.Parse()
	if len(p.Errors()) > 0 {
		t.Fatal(p.Errors())
	}

	env := NewEnvironment()
	env.Set("a", int64(1))
	env.Set("b", int64(4))
	env.Set("c", int64(3))
	env.Set("d", "hello")

	res := Match(exp, env)
	if !res {
		t.Fail()
	}
}

func BenchmarkVisitor(b *testing.B) {
	input := `a==1 && b==2 || c in [3]`
	l := lexer.New(input)
	p := parser.New(l)
	exp := p.Parse()
	if len(p.Errors()) > 0 {
		b.Fatal(p.Errors())
	}

	env := NewEnvironment()
	env.Set("a", int64(1))
	env.Set("b", int64(2))
	env.Set("c", int64(3))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := Match(exp, env)
		if !res {
			b.Fail()
		}
	}
}
