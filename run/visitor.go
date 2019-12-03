package run

import (
	"github.com/MoonShining/condition/ast"
	"github.com/MoonShining/condition/token"
)

func Match(exp ast.Expression, env *Environment) bool {
	res := run(exp, env)
	switch v := res.(type) {
	case bool:
		return v
	default:
		return false
	}
}

func run(exp ast.Expression, env *Environment) interface{} {
	switch e := exp.(type) {
	case *ast.InfixExpression:
		left := run(e.Left, env)
		right := run(e.Right, env)

		switch e.Token.Type {
		case token.EQ:
			return left == right
		case token.NOT_EQ:
			return left != right
		case token.LT:
			return binaryOp(left, right, func(l int64, r int64) bool {
				return l < r
			}, func(l string, r string) bool {
				return l < r
			})
		case token.LTE:
			return binaryOp(left, right, func(l int64, r int64) bool {
				return l <= r
			}, func(l string, r string) bool {
				return l <= r
			})
		case token.GT:
			return binaryOp(left, right, func(l int64, r int64) bool {
				return l > r
			}, func(l string, r string) bool {
				return l > r
			})
		case token.GTE:
			return binaryOp(left, right, func(l int64, r int64) bool {
				return l >= r
			}, func(l string, r string) bool {
				return l >= r
			})
		case token.IN:
			return inOp(left, right)
		case token.AND:
			return left.(bool) && right.(bool)
		case token.OR:
			return left.(bool) || right.(bool)
		default:
			return false
		}
	case *ast.Identifier:
		id, ok := env.Get(e.Value)
		if ok {
			return id
		} else {
			return nil
		}
	case *ast.StringLiteral:
		return e.Value
	case *ast.IntegerLiteral:
		return e.Value
	case *ast.Null:
		return nil
	case *ast.Boolean:
		return e.Value
	case *ast.ArrayLiteral:
		result := make([]interface{}, len(e.Elements))
		for i, e := range e.Elements {
			value := run(e, env)
			result[i] = value
		}
		return result
	default:
		return false
	}
}

type intFunc = func(left int64, right int64) bool
type strFunc = func(left string, right string) bool

func binaryOp(left, right interface{}, f1 intFunc, f2 strFunc) bool {
	switch l := left.(type) {
	case int64:
		r, ok := right.(int64)
		if !ok {
			return false
		}
		return f1(l, r)
	case string:
		r, ok := right.(string)
		if !ok {
			return false
		}
		return f2(l, r)
	default:
		return false
	}
}

func inOp(left, right interface{}) bool {
	arr, ok := right.([]interface{})
	if !ok {
		return false
	}

	res := false

	for _, ele := range arr {
		if left == ele {
			res = true
			break
		}
	}
	return res
}
