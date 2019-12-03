package condition

import (
	"fmt"
	"github.com/MoonShining/condition/ast"
	"github.com/MoonShining/condition/lexer"
	"github.com/MoonShining/condition/parser"
	"github.com/MoonShining/condition/run"
)

type Condition struct {
	expr ast.Expression
}

func NewCondition(content string) (*Condition, error) {
	l := lexer.New(content)
	p := parser.New(l)
	expr := p.Parse()
	if len(p.Errors()) > 0 {
		return nil, fmt.Errorf("%v", p.Errors())
	}

	return &Condition{expr: expr}, nil
}

func (c *Condition) Match(env *run.Environment) bool {
	return run.Match(c.expr, env)
}
