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

// influunt_NewGraph returns a new influunt graph as void pointer
//export influunt_NewGraph
func influunt_NewGraph(self *pyObject, args *pyObject) *pyObject {
	return pointerToCapsule(pointer.Save(influunt.NewGraph()))
}

// influunt_GraphNodeByName returns a graph node by its name and index
//export influunt_GraphNodeByName
func influunt_GraphNodeByName(self *pyObject, args *pyObject) *pyObject {
	graphCapsule, name, index := parse3ObjectFromArgs(args)
	graph := graphFromPointer(capsuleToPointer(graphCapsule))

	nameInterface, err := convertPyObjectToInterface(name)
	if err != nil {
		panic(err)
	}

	nameString, ok := nameInterface.(string)
	if !ok {
		panic(errors.New("name should be a string"))
	}

	indexInterface, err := convertPyObjectToInterface(index)
	if err != nil {
		panic(err)
	}

	indexInt, ok := indexInterface.(int)
	if !ok {
		panic(errors.New("index should be an int"))
	}

	node := graph.NodeByName(nameString, indexInt)
	if node == nil {
		return C.Py_None
	}

	return pointerToCapsule(nodeToPointer(*node))
}

// influunt_ReadGraphFromFile loads graph for a given filepath
//export influunt_ReadGraphFromFile
func influunt_ReadGraphFromFile(self *pyObject, filepath *pyObject) *pyObject {
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

	graph, err := influunt.ReadGraph(file)
	if err != nil {
		panic(err)
	}

	return pointerToCapsule(pointer.Save(graph))
}

// influunt_WriteGraphToFile writes graph to a given filepath
//export influunt_WriteGraphToFile
func influunt_WriteGraphToFile(self *pyObject, args *pyObject) *pyObject {
	graphCapsule, filepath := parse2ObjectFromArgs(args)

	graph := graphFromPointer(capsuleToPointer(graphCapsule))

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

	err = influunt.WriteGraph(graph, file)
	if err != nil {
		panic(err)
	}

	return C.Py_None
}
