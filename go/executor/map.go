package executor

import (
	"reflect"

	"github.com/orktes/influunt/go"
	"github.com/pkg/errors"
)

func mapOp(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
	if len(inputs) != 2 {
		return nil, errors.New("operation takes two arguments")
	}

	iterables, err := e.executeOp(c, inputs[0])
	if err != nil {
		return nil, err
	}

	s := reflect.ValueOf(iterables)
	if attrs == nil {
		return nil, errors.New("not attributes defined")
	}

	argument, ok := attrs["argument"]
	if !ok {
		return nil, errors.New("argument not defined in attributes")
	}

	argumentNode, ok := argument.(influunt.Node)
	if !ok {
		return nil, errors.New("arguments is not a node")
	}

	index, ok := attrs["index"]
	if !ok {
		return nil, errors.New("index not defined in attributes")
	}

	indexNode, ok := index.(influunt.Node)
	if !ok {
		return nil, errors.New("index is not a node")
	}

	var results reflect.Value

	length := s.Len()

	for i := 0; i < length; i++ {
		value := s.Index(i)
		// TODO: memory pooling
		subContext := c.SubContext(map[influunt.Node]interface{}{
			argumentNode: value.Interface(),
			indexNode:    i,
		})
		res, err := e.executeOp(subContext, inputs[1])
		if err != nil {
			return nil, err
		}

		if i == 0 {
			results = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(res)), length, length)
		}

		v := results.Index(i)
		v.Set(reflect.ValueOf(res))
	}

	return []interface{}{results.Interface()}, nil
}

func init() {
	AddOperation("Map", mapOp)
}
