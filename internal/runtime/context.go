package runtime

import "fmt"

type Context struct {
	values map[string]ConcreteValue
	parent *Context
}

func NewContext() *Context {
	return &Context{
		values: make(map[string]ConcreteValue),
	}
}

func (c *Context) SetValue(name string, v ConcreteValue) {
	c.values[name] = v
}

func (c *Context) GetValue(name string) ConcreteValue {
	if v, ok := c.values[name]; ok {
		return v
	}
	if c.parent != nil {
		return c.parent.GetValue(name)
	}
	panic(fmt.Errorf("trying to get undefined value from context: %s", name))
}

func (c *Context) HasValue(name string) bool {
	_, ok := c.values[name]
	return ok
}
