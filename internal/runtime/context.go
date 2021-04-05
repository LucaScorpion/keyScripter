package runtime

import "fmt"

type Context struct {
	values map[string]Value
}

func NewContext() *Context {
	return &Context{
		values: make(map[string]Value),
	}
}

func (c *Context) SetValue(name string, v Value) {
	if va, ok := v.(VariableValue); ok {
		panic(fmt.Errorf("context can only store concrete values, tried to store variable: %s", va.ref))
	}

	c.values[name] = v
}

func (c *Context) GetValue(name string) Value {
	return c.values[name]
}
