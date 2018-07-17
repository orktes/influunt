package executor

import "github.com/orktes/influunt/go"

// Context execution context for run
type Context struct {
	node          influunt.Node
	inputs        map[influunt.Node]interface{}
	vals          [][]interface{}
	parentContext *Context
}

// SubContext returns a new subcontext
func (c *Context) SubContext(inputs map[influunt.Node]interface{}) *Context {
	return &Context{
		inputs:        inputs,
		vals:          c.vals,
		parentContext: c,
	}
}

// GetInput returns an input argument from context
func (c *Context) GetInput(node influunt.Node) (interface{}, bool) {
	if c.inputs != nil {
		val, ok := c.inputs[node]
		if !ok && c.parentContext != nil {
			return c.parentContext.GetInput(node)
		}

		return val, true
	}

	return nil, false
}
