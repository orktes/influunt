package executor

import (
	"github.com/orktes/influunt/go"
	"github.com/pkg/errors"
)

func getAttrOp(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
	if len(inputs) != 2 {
		return nil, errors.New("operation takes two arguments")
	}

	m, err := e.ExecuteOp(c, inputs[0])
	if err != nil {
		return nil, err
	}

	key, err := e.ExecuteOp(c, inputs[1])
	if err != nil {
		return nil, err
	}

	strKey, ok := key.(string)
	if !ok {
		return nil, errors.New("key should be a string")
	}

	// TODO support other types also
	mapVal, ok := m.(map[string]interface{})
	if !ok {
		return nil, errors.New("not a map")
	}

	return []interface{}{
		mapVal[strKey],
	}, nil
}

func init() {
	AddOperation("GetAttr", getAttrOp)
}
