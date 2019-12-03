package run

func NewEnvironment() *Environment {
	s := make(map[string]interface{})
	return &Environment{store: s}
}

type Environment struct {
	store map[string]interface{}
}

func (e *Environment) Get(name string) (interface{}, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val interface{}) interface{} {
	e.store[name] = val
	return val
}
