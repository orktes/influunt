package executor

import (
	"github.com/orktes/influunt/go"
	"github.com/pkg/errors"
)

func constOp(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
	if attrs == nil {
		return nil, errors.New("not attributes defined")
	}

	val, ok := attrs["value"]
	if !ok {
		return nil, errors.New("value not defined in attributes")
	}

	return []interface{}{
		val,
	}, nil
}

func init() {
	AddOperation("Const", constOp)
}
