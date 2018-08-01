package executor

import (
	"github.com/orktes/influunt/go"
	"github.com/pkg/errors"
)

// ErrMissingInput is returned if given input map is missing inputs
var ErrMissingInput = errors.New("missing input variable")

// ModelExecutor is used to execute models (graph + signature)
type ModelExecutor struct {
	model    *influunt.Model
	executor *Executor

	outputs     []influunt.Node
	outputNames []string
}

// NewModelExecutor returns a new model executor for a given model
func NewModelExecutor(m *influunt.Model) (*ModelExecutor, error) {
	e, err := NewExecutor(m.Graph)
	if err != nil {
		return nil, err
	}

	me := &ModelExecutor{model: m, executor: e}

	return me, me.init()
}

func (me *ModelExecutor) init() error {
	me.outputs = make([]influunt.Node, 0, len(me.model.Outputs))
	me.outputNames = make([]string, 0, len(me.model.Outputs))

	for name, node := range me.model.Outputs {
		me.outputs = append(me.outputs, node)
		me.outputNames = append(me.outputNames, name)
	}

	return nil
}

// Run executes graph with given input map and output filters
func (me *ModelExecutor) Run(inputs map[string]interface{}) (map[string]interface{}, error) {
	executorInputs := make(map[influunt.Node]interface{}, len(inputs))

	for name, node := range me.model.Inputs {
		input, ok := inputs[name]
		if !ok {
			return nil, errors.Wrapf(ErrMissingInput, "missing input variable %s", name)
		}

		executorInputs[node] = input
	}

	results, err := me.executor.Run(executorInputs, me.outputs)
	if err != nil {
		return nil, err
	}

	outputs := make(map[string]interface{}, len(me.outputs))
	for i, name := range me.outputNames {
		outputs[name] = results[i]
	}

	return outputs, nil
}
