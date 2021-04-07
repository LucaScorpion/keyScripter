package runtime

type Context struct {
	values map[string]ConcreteValue
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
	return c.values[name]
}

func (c *Context) HasValue(name string) bool {
	_, ok := c.values[name]
	return ok
}
