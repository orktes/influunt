package main

// #cgo pkg-config: python-3.6
// #include <Python.h>
import "C"
import (
	"errors"

	"github.com/mattn/go-pointer"
	"github.com/orktes/influunt/go/executor"
)

// influunt_NewModelExecutor returns a new executor
//export influunt_NewModelExecutor
func influunt_NewModelExecutor(self, mCapsule *pyObject) *pyObject {
	m := capsuleToPointer(mCapsule)
	model := modelFromPointer(m)
	executor, err := executor.NewModelExecutor(model)
	if err != nil {
		panic(err)
	}
	return pointerToCapsule(pointer.Save(executor))
}

// influunt_ModelExecutorRun runs model executor executor
//export influunt_ModelExecutorRun
func influunt_ModelExecutorRun(self, args *pyObject) *C.PyObject {
	eCapsule, inputs := parse2ObjectFromArgs(args)
	e := capsuleToPointer(eCapsule)
	exec := pointer.Restore(e).(*executor.ModelExecutor)

	inputInterface, err := convertPyObjectToInterface(inputs)
	if err != nil {
		panic(err)
	}

	inputMap, ok := inputInterface.(map[string]interface{})
	if !ok {
		panic(errors.New("inputs should be a dictionary"))
	}

	res, err := exec.Run(inputMap)
	if err != nil {
		panic(err)
	}

	resPyObj, err := convertGoTypeToPyObject(res)
	if err != nil {
		panic(err)
	}
	return resPyObj
}

// influunt_ModelExecutorRunAsync runs executor async
//export influunt_ModelExecutorRunAsync
func influunt_ModelExecutorRunAsync(self, args *pyObject) *C.PyObject {
	eCapsule, inputs, callback := parse3ObjectFromArgs(args)
	e := capsuleToPointer(eCapsule)
	exec := pointer.Restore(e).(*executor.ModelExecutor)

	inputInterface, err := convertPyObjectToInterface(inputs)
	if err != nil {
		panic(err)
	}

	inputMap, ok := inputInterface.(map[string]interface{})
	if !ok {
		panic(errors.New("inputs should be a dictionary"))
	}
	go func() {
		result, err := exec.Run(inputMap)
		if err != nil {
			panic(err)
		}

		gstate := C.PyGILState_Ensure()
		defer C.PyGILState_Release(gstate)

		resPyObj, err := convertGoTypeToPyObject(result)
		if err != nil {
			panic(err)
		}

		args := C.PyTuple_New(C.long(1))
		C.PyTuple_SetItem(args, C.long(0), resPyObj)

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
