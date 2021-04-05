package runtime

import "fmt"

type context struct {
	values map[string]Value
}

func (c *context) setValue(name string, v Value) {
	if va, ok := v.(VariableValue); ok {
		panic(fmt.Errorf("context can only store concrete values, tried to store variable: %s", va.ref))
	}

	c.values[name] = v
}

func (c *context) getValue(name string) Value {
	return c.values[name]
}
