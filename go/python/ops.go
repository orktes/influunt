package main

// #cgo pkg-config: python-3.6
// #include <Python.h>
import "C"
import (
	"github.com/orktes/influunt/go"
	"errors"
	"unsafe"
)

// influunt_OpMap maps over a given list
//export influunt_OpMap
func influunt_OpMap(self, args *pyObject) *pyObject {
	gCapsule, listCapsule, mapFn := parse3ObjectFromArgs(args)
	g := graphFromPointer(capsuleToPointer(gCapsule))
	list := nodeFromPointer(capsuleToPointer(listCapsule))

	return pointerToCapsule(nodeToPointer(influunt.Map(g, list, func(item influunt.Node, index influunt.Node) influunt.Node {
		args := C.PyTuple_New(C.long(2))
		defer pyRelease(args)

		pyItem := pointerToCapsule(nodeToPointer(item))
		pyIndex := pointerToCapsule(nodeToPointer(index))

		C.PyTuple_SetItem(args, C.long(0), pyItem)
		C.PyTuple_SetItem(args, C.long(1), pyIndex)

		res := C.PyObject_CallObject(mapFn, args)
		if res == nil {
			C.PyErr_Print()
		}

		return nodeFromPointer(capsuleToPointer(res))
	})))
}

// influunt_GraphAddOp adds a new operation to the graph based on the given job spec
//export influunt_GraphAddOp
func influunt_GraphAddOp(self, args *pyObject) *pyObject {
	graphCapsule, spec := parse2ObjectFromArgs(args)
	graph := graphFromPointer(capsuleToPointer(graphCapsule))
	specMapInterface, err := convertPyObjectToInterface(spec)
	if err != nil {
		panic(err)
	}

	specMap, ok := specMapInterface.(map[string]interface{})
	if !ok {
		panic(errors.New("OpSpec is not a dict"))
	}

	opSpec := influunt.OpSpec{}
	resLength := 1

	if typ, ok := specMap["type"]; ok {
		if typ, ok := typ.(string); ok {
			opSpec.Type = typ
		}
	}

	if attrs, ok := specMap["attrs"]; ok {
		if attrs, ok := attrs.(map[string]interface{}); ok {
			opSpec.Attrs = attrs
		}
	}

	if inputs, ok := specMap["inputs"]; ok {
		if inputs, ok := inputs.([]unsafe.Pointer); ok {
			for _, ptr := range inputs {
				opSpec.Inputs = append(opSpec.Inputs, nodeFromPointer(ptr))
			}
			
		}
	}

	if results, ok := specMap["results"]; ok {
		if results, ok := results.(int); ok {
			resLength = results
		}
	}

	op := graph.AddOperation(opSpec)

	nodes := make([]unsafe.Pointer, resLength)
	for i := 0; i < resLength; i++ {
		nodes[i] = nodeToPointer(op.Output(i))
	}

	pyRes, err := convertGoTypeToPyObject(nodes)
	if err != nil {
		panic(err)
	}
	
	return pyRes
}
