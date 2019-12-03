package condition

import (
	"github.com/MoonShining/condition/run"
	"testing"
)

func TestCondition_Match(t *testing.T) {
	c, _ := NewCondition(`(false || a==1) && (b==2||c==3) && (d in ["hello"])`)
	env := run.NewEnvironment()
	env.Set("a", int64(1))
	env.Set("b", int64(4))
	env.Set("c", int64(3))
	env.Set("d", "hello")
	if !c.Match(env) {
		t.Fail()
	}
}
