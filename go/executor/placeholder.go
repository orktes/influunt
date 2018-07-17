package executor

import (
	"github.com/orktes/influunt/go"
)

func placeholderOp(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
	return []interface{}{
		c.GetInput(c.node),
	}, nil
}

func init() {
	AddOperation("Placeholder", placeholderOp)
}
