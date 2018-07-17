package executor

import "github.com/orktes/influunt/go"

// Op function signature for graph operations
type Op func(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error)

var registeredOps = map[string]Op{}

// AddOperation adds a new operation to be used for execution
func AddOperation(opType string, op Op) {
	registeredOps[opType] = op
}
