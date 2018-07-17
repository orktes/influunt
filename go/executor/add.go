package executor

import (
	"fmt"

	"github.com/orktes/influunt/go"
	"github.com/pkg/errors"
)

func addOpInt(a int, b interface{}) (interface{}, error) {
	switch b := b.(type) {
	case int:
		return a + b, nil
	case float64:
		return float64(a) + b, nil
	default:
		return nil, fmt.Errorf("unsupported type on the right side %T", b)
	}
}

func addOpFloat64(a float64, b interface{}) (interface{}, error) {
	switch b := b.(type) {
	case int:
		return a + float64(b), nil
	case float64:
		return a + b, nil
	default:
		return nil, fmt.Errorf("unsupported type on the right side %T", b)
	}
}

func addOpString(a string, b interface{}) (interface{}, error) {
	switch b := b.(type) {
	case string:
		return a + b, nil
	default:
		return nil, fmt.Errorf("unsupported type on the right side %T", b)
	}
}

func addOp(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
	if len(inputs) != 2 {
		return nil, errors.New("operation takes two arguments")
	}

	var res interface{}

	a, err := e.executeOp(c, inputs[0])
	if err != nil {
		return nil, err
	}
	b, err := e.executeOp(c, inputs[1])
	if err != nil {
		return nil, err
	}

	switch a := a.(type) {
	case int:
		res, err = addOpInt(a, b)
	case float64:
		res, err = addOpFloat64(a, b)
	case string:
		res, err = addOpString(a, b)
	}

	return []interface{}{res}, err
}

func init() {
	AddOperation("Add", addOp)
}
