package runtime

import "fmt"

type Context struct {
	values map[string]*Value
	parent *Context
}

func NewContext(parent *Context) *Context {
	return &Context{
		values: map[string]*Value{},
		parent: parent,
	}
}

func RootContext() *Context {
	ctx := NewContext(nil)

	// Store all the native functions in the context.
	for n, f := range nativeFunctions {
		ctx.SetValue(n, f)
	}

	return ctx
}

func (c *Context) Parent() *Context {
	return c.parent
}

func (c *Context) SetValue(name string, v *Value) {
	c.values[name] = v
}

func (c *Context) GetValue(name string) *Value {
	if v, ok := c.values[name]; ok {
		return v
	}
	if c.parent != nil {
		return c.parent.GetValue(name)
	}
	panic(fmt.Errorf("trying to get undefined value from context: %s", name))
}

func (c *Context) HasValue(name string) bool {
	if _, ok := c.values[name]; ok {
		return ok
	}
	if c.parent != nil {
		return c.parent.HasValue(name)
	}
	return false
}
