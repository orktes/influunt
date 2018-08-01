package executor

import (
	"github.com/orktes/influunt/go"
	"github.com/pkg/errors"
)

func condOp(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
	if len(inputs) != 3 {
		return nil, errors.New("wrong number of arguments")
	}

	pred, err := e.ExecuteOp(c, inputs[0])
	if err != nil {
		return nil, err
	}

	var res interface{}
	if pred == true {
		res, err = e.ExecuteOp(c, inputs[1])
	} else {
		res, err = e.ExecuteOp(c, inputs[2])
	}

	if err != nil {
		return nil, err
	}

	return []interface{}{res}, nil
}

func init() {
	AddOperation("Cond", condOp)
}
