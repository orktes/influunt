package executor

import (
	"github.com/orktes/influunt/go"
	"github.com/pkg/errors"
)

func placeholderOp(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
	val, ok := c.GetInput(c.node)
	if !ok {
		return nil, errors.New("input value missing")
	}

	return []interface{}{
		val,
	}, nil
}

func init() {
	AddOperation("Placeholder", placeholderOp)
}
