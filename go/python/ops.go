package main

// #cgo pkg-config: python-3.6
// #include <Python.h>
import "C"
import (
	"github.com/orktes/influunt/go"
)

// influunt_OpAdd operations adds two nodes together (a + b)
//export influunt_OpAdd
func influunt_OpAdd(self, args *pyObject) *pyObject {
	g, a, b := parse3PointerFromArgs(args)
	return pointerToCapsule(
		nodeToPointer(
			influunt.Add(graphFromPointer(g), nodeFromPointer(a), nodeFromPointer(b)),
		),
	)
}

// influunt_OpPlaceholder returns a new placeholder node
//export influunt_OpPlaceholder
func influunt_OpPlaceholder(self, args *pyObject) *pyObject {
	g := parse1PointerFromArgs(args)
	return pointerToCapsule(nodeToPointer(influunt.Placeholder(graphFromPointer(g))))
}

// influunt_OpConst returns a constant value node
//export influunt_OpConst
func influunt_OpConst(self, args *pyObject) *pyObject {
	gCapsule, valObject := parse2ObjectFromArgs(args)
	g := capsuleToPointer(gCapsule)
	v, err := convertPyObjectToInterface(valObject)
	if err != nil {
		panic(err)
	}
	return pointerToCapsule(nodeToPointer(influunt.Const(graphFromPointer(g), v)))
}

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

// influunt_OpParseJSON operation parses json
//export influunt_OpParseJSON
func influunt_OpParseJSON(self, args *pyObject) *pyObject {
	g, json := parse2PointerFromArgs(args)
	return pointerToCapsule(
		nodeToPointer(
			influunt.ParseJSON(graphFromPointer(g), nodeFromPointer(json)),
		),
	)
}

// influunt_OpGetAttr returns an attribute from an object or map
//export influunt_OpGetAttr
func influunt_OpGetAttr(self, args *pyObject) *pyObject {
	g, m, key := parse3PointerFromArgs(args)
	return pointerToCapsule(
		nodeToPointer(
			influunt.GetAttr(graphFromPointer(g), nodeFromPointer(m), nodeFromPointer(key)),
		),
	)
}

// influunt_OpCond returs a if pred is "true" and b if pred is "false"
//export influunt_OpCond
func influunt_OpCond(self, args *pyObject) *pyObject {
	g, pred, a, b := parse4PointerFromArgs(args)
	return pointerToCapsule(
		nodeToPointer(
			graphFromPointer(g).AddOperation(influunt.OpSpec{
				Type:   "Cond",
				Inputs: []influunt.Node{nodeFromPointer(pred), nodeFromPointer(a), nodeFromPointer(b)},
			}).Output(0),
		),
	)
}