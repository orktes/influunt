package executor

import (
	"github.com/pkg/errors"

	"github.com/orktes/influunt/go"
)

// ErrOpNotRegistered error returned when Operation defined in the graph is not registered
var ErrOpNotRegistered = errors.New("Operation not registered")

// Executor executes given graph
type Executor struct {
	g   *influunt.Graph
	ops []Op
}

// NewExecutor returns a new executor
func NewExecutor(g *influunt.Graph) (*Executor, error) {
	e := &Executor{g: g}
	return e, e.init()
}

func (e *Executor) init() error {
	ops := make([]Op, len(e.g.Operations))
	for i, op := range e.g.Operations {
		opImpl, ok := registeredOps[op.Spec.Type]
		if !ok {
			return errors.Wrapf(ErrOpNotRegistered, "operation: %s", op.Spec.Type)
		}
		ops[i] = opImpl
	}
	e.ops = ops
	return nil
}

func (e *Executor) executeOp(context *Context, n influunt.Node) (interface{}, error) {
	id := n.ID

	vals := context.vals
	val := vals[id[0]]
	if val != nil {
		return val[id[1]], nil
	}

	operation := e.g.Operations[id[0]]
	spec := operation.Spec

	var err error

	opFunc := e.ops[id[0]]

	context.node = n
	val, err = opFunc(context, e, spec.Inputs, spec.Attrs)
	if err != nil {
		return nil, err
	}

	if operation.Cache {
		vals[id[0]] = val
	}
	return val[id[1]], err
}

// Run executes graph with given input map and output filters
func (e *Executor) Run(inputs map[influunt.Node]interface{}, outputs []influunt.Node) ([]interface{}, error) {
	// TODO: Use memory pooling for this
	vals := make([][]interface{}, len(e.g.Operations))
	context := &Context{
		inputs: inputs,
		vals:   vals,
	}

	var err error
	results := make([]interface{}, len(outputs))
	for i, output := range outputs {
		results[i], err = e.executeOp(context, output)
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}
