package main

// #cgo pkg-config: python-3.6
// #include <Python.h>
import "C"
import (
	"errors"
	"os"

	"github.com/mattn/go-pointer"
	"github.com/orktes/influunt/go"
)

// influunt_ReadModelFromFile loads graph for a given filepath
//export influunt_ReadModelFromFile
func influunt_ReadModelFromFile(self *pyObject, filepath *pyObject) *pyObject {
	val, err := convertPyObjectToInterface(filepath)
	if err != nil {
		panic(err)
	}

	fpath, ok := val.(string)
	if !ok {
		panic(errors.New("filepath should be a string"))
	}

	file, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	model, err := influunt.ReadModel(file)
	if err != nil {
		panic(err)
	}

	return pointerToCapsule(pointer.Save(model))
}

// influunt_WriteModelToFile writes model to a given filepath
//export influunt_WriteModelToFile
func influunt_WriteModelToFile(self *pyObject, args *pyObject) *pyObject {
	graphCapsule, inputsDict, outputsDict, filepath := parse4ObjectFromArgs(args)
	graph := graphFromPointer(capsuleToPointer(graphCapsule))

	inputs, err := convertPyDictToSignature(inputsDict)
	if err != nil {
		panic(err)
	}

	outputs, err := convertPyDictToSignature(outputsDict)
	if err != nil {
		panic(err)
	}

	val, err := convertPyObjectToInterface(filepath)
	if err != nil {
		panic(err)
	}

	fpath, ok := val.(string)
	if !ok {
		panic(errors.New("filepath should be a string"))
	}

	file, err := os.Create(fpath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = influunt.WriteModel(&influunt.Model{
		Graph:   graph,
		Inputs:  inputs,
		Outputs: outputs,
	}, file)
	if err != nil {
		panic(err)
	}

	return C.Py_None
}

func convertPyDictToSignature(obj *pyObject) (map[string]influunt.Node, error) {
	m := map[string]influunt.Node{}

	keys := C.PyDict_Keys(obj)
	length := int(C.PyList_Size(keys))

	for i := 0; i < length; i++ {
		key := C.PyList_GetItem(keys, C.long(i))
		keyVal, err := convertPyObjectToInterface(key)

		val := C.PyDict_GetItem(obj, key)
		node := nodeFromPointer(capsuleToPointer(val))

		if err != nil {
			return nil, err
		}

		m[keyVal.(string)] = node

	}

	return m, nil
}
