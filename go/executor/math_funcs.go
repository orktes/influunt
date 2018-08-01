package executor

import (
	"fmt"
	"reflect"
)

type mathFuncInfo struct {
	aType reflect.Type
	bType reflect.Type
}

var mathFuncs = map[string]map[mathFuncInfo]reflect.Value{}

func getMathFunc(op string) func(a, b interface{}) (interface{}, error) {
	funcMap, ok := mathFuncs[op]
	if !ok {
		funcMap = map[mathFuncInfo]reflect.Value{}
		mathFuncs[op] = funcMap
	}

	return func(a, b interface{}) (interface{}, error) {
		aType := reflect.TypeOf(a)
		bType := reflect.TypeOf(b)

		aValue := reflect.ValueOf(a)
		bValue := reflect.ValueOf(b)

		fn, ok := funcMap[mathFuncInfo{
			aType: aType,
			bType: bType,
		}]

		if !ok {
			return nil, fmt.Errorf("no %s op for %T and %T", op, a, b)
		}

		res := fn.Call([]reflect.Value{aValue, bValue})

		return res[0].Interface(), nil
	}
}

func addMathFunc(op string, fn interface{}) bool {
	value := reflect.ValueOf(fn)
	typ := reflect.TypeOf(fn)

	info := mathFuncInfo{
		aType: typ.In(0),
		bType: typ.In(1),
	}

	funcMap, ok := mathFuncs[op]
	if !ok {
		funcMap = map[mathFuncInfo]reflect.Value{}
	}

	funcMap[info] = value
	mathFuncs[op] = funcMap
	return true
}
