package executor

import (
	"encoding/json"
	"fmt"

	"github.com/orktes/influunt/go"
	"github.com/pkg/errors"
)

func parseJSONOp(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
	if len(inputs) == 0 {
		return nil, errors.New("operation takes one argument")
	}

	str, err := e.executeOp(c, inputs[0])
	if err != nil {
		return nil, err
	}

	var byteData []byte
	switch v := str.(type) {
	case string:
		byteData = []byte(v)
	case []byte:
		byteData = v
	default:
		return nil, fmt.Errorf("unsupported data type %T", v)
	}

	m := map[string]interface{}{}
	err = json.Unmarshal(byteData, &m)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		m,
	}, nil
}

func init() {
	AddOperation("ParseJSON", parseJSONOp)
}
