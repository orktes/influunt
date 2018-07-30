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
