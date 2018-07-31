package main

// #cgo pkg-config: python-3.6
// #include <Python.h>
import "C"

// influunt_NodeGetName returns name of the node
//export influunt_NodeGetName
func influunt_NodeGetName(self *pyObject, args *pyObject) *pyObject {
	node := nodeFromPointer(capsuleToPointer(args))
	name, _ := convertGoTypeToPyObject(node.Name)
	return name
}
