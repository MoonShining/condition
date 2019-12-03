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

func TestCondition_MatchEQ(t *testing.T) {
	c, _ := NewCondition(`a == 1`)
	env := run.NewEnvironment()
	env.Set("a", int64(1))
	if !c.Match(env) {
		t.Fail()
	}
}

func TestCondition_MatchIn(t *testing.T) {
	c, _ := NewCondition(`a in ["1", 1]`)
	env := run.NewEnvironment()
	env.Set("a", int64(1))
	if !c.Match(env) {
		t.Fail()
	}
	env.Set("a", "2")
	if c.Match(env) {
		t.Fail()
	}
}

func TestCondition_MatchGtLtGteLte(t *testing.T) {
	c, _ := NewCondition(`a>1 && b<2 && c>=4 && d<=5`)
	env := run.NewEnvironment()
	env.Set("a", int64(2))
	env.Set("b", int64(1))
	env.Set("c", int64(4))
	env.Set("d", int64(5))
	if !c.Match(env) {
		t.Fail()
	}
}

func TestCondition_MatchParen(t *testing.T) {
	c, _ := NewCondition(`(a=="hello" && b==1) || true`)
	env := run.NewEnvironment()
	env.Set("a", int64(2))
	env.Set("b", int64(1))
	if !c.Match(env) {
		t.Fail()
	}
}
