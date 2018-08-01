package main

// #cgo pkg-config: python-3.6
// #include <Python.h>
import "C"
import (
	"github.com/mattn/go-pointer"
	"github.com/orktes/influunt/go"
	"github.com/orktes/influunt/go/executor"
)

// influunt_NewExecutor returns a new executor
//export influunt_NewExecutor
func influunt_NewExecutor(self, gCapsule *pyObject) *pyObject {
	g := capsuleToPointer(gCapsule)
	graph := graphFromPointer(g)
	executor, err := executor.NewExecutor(graph)
	if err != nil {
		panic(err)
	}
	return pointerToCapsule(pointer.Save(executor))
}

// influunt_ExecutorAddOperation adds a new operation
//export influunt_ExecutorAddOperation
func influunt_ExecutorAddOperation(self, args *pyObject) *C.PyObject {
	name, fn := parse2ObjectFromArgs(args)
	pyRetain(fn)

	nameStr, err := convertPyObjectToInterface(name)
	if err != nil {
		panic(err)
	}

	executor.AddOperation(nameStr.(string), func(c *executor.Context, e *executor.Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
		args := C.PyTuple_New(C.long(len(inputs)))
		defer pyRelease(args)

		for i, input := range inputs {
			val, err := e.ExecuteOp(c, input)
			if err != nil {
				return nil, err
			}

			pyVal, err := convertGoTypeToPyObject(val)
			if err != nil {
				return nil, err
			}

			C.PyTuple_SetItem(args, C.long(i), pyVal)
		}

		res := C.PyObject_CallObject(fn, args)
		if res == nil {
			// TODO return error
			C.PyErr_Print()
		}

		resGoVal, err := convertPyObjectToInterface(res)
		if err != nil {
			return nil, err
		}

		return []interface{}{resGoVal}, nil
	})

	pyRetain(C.Py_None)
	return C.Py_None
}

// influunt_ExecutorRun runs executor
//export influunt_ExecutorRun
func influunt_ExecutorRun(self, args *pyObject) *C.PyObject {
	eCapsule, inputs, outputs := parse3ObjectFromArgs(args)
	e := capsuleToPointer(eCapsule)
	exec := pointer.Restore(e).(*executor.Executor)

	inputMap, err := convertPyDictNodeMap(inputs)
	if err != nil {
		panic(err)
	}

	outputArr, err := convertPyListToNodeArr(outputs)
	if err != nil {
		panic(err)
	}

	res, err := exec.Run(inputMap, outputArr)
	if err != nil {
		panic(err)
	}

	resPyObj, err := convertGoTypeToPyObject(res)
	if err != nil {
		panic(err)
	}
	return resPyObj
}

// influunt_ExecutorRunAsync runs executor async
//export influunt_ExecutorRunAsync
func influunt_ExecutorRunAsync(self, args *pyObject) *C.PyObject {
	eCapsule, inputs, outputs, callback := parse4ObjectFromArgs(args)
	e := capsuleToPointer(eCapsule)
	exec := pointer.Restore(e).(*executor.Executor)

	inputMap, err := convertPyDictNodeMap(inputs)
	if err != nil {
		panic(err)
	}

	outputArr, err := convertPyListToNodeArr(outputs)
	if err != nil {
		panic(err)
	}
	go func() {
		responses, err := exec.Run(inputMap, outputArr)
		if err != nil {
			panic(err)
		}

		gstate := C.PyGILState_Ensure()
		defer C.PyGILState_Release(gstate)

		args := C.PyTuple_New(C.long(len(responses)))
		for i, val := range responses {
			pyVal, err := convertGoTypeToPyObject(val)
			if err != nil {
				panic(err)
			}

			C.PyTuple_SetItem(args, C.long(i), pyVal)
		}

		res := C.PyObject_CallObject(callback, args)
		if res == nil {
			// TODO handle error
			C.PyErr_Print()
		}
		pyRelease(args)
	}()

	pyRetain(C.Py_None)
	return C.Py_None
}

func convertPyListToNodeArr(list *pyObject) ([]influunt.Node, error) {
	length := int(C.PyList_Size(list))
	arr := make([]influunt.Node, 0, length)

	for i := 0; i < length; i++ {
		item := C.PyList_GetItem(list, C.long(i))
		node := nodeFromPointer(capsuleToPointer(item))
		arr = append(arr, node)
	}

	return arr, nil
}

func convertPyDictNodeMap(obj *pyObject) (map[influunt.Node]interface{}, error) {
	m := map[influunt.Node]interface{}{}

	keys := C.PyDict_Keys(obj)
	length := int(C.PyList_Size(keys))

	for i := 0; i < length; i++ {
		key := C.PyList_GetItem(keys, C.long(i))
		node := nodeFromPointer(capsuleToPointer(key))

		val := C.PyDict_GetItem(obj, key)
		valVal, err := convertPyObjectToInterface(val)
		if err != nil {
			return nil, err
		}

		m[node] = valVal

	}

	return m, nil
}
