package main

// #cgo pkg-config: python-3.6
// #include <Python.h>
// void influunt_destructCapsule(PyObject* p0);
//
// PyObject* influunt_pointerToCapsule(void* ptr)
// {
//		return PyCapsule_New(ptr, "influunt", influunt_destructCapsule);
// }
import "C"
import (
	"unsafe"

	"github.com/mattn/go-pointer"
	"github.com/orktes/influunt/go"
)

func parse1PointerFromArgs(args *C.PyObject) unsafe.Pointer {
	a := parse1ObjectFromArgs(args)
	return capsuleToPointer(a)
}

func parse2PointerFromArgs(args *C.PyObject) (unsafe.Pointer, unsafe.Pointer) {
	a, b := parse2ObjectFromArgs(args)
	return capsuleToPointer(a), capsuleToPointer(b)
}

func parse3PointerFromArgs(args *C.PyObject) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	a, b, c := parse3ObjectFromArgs(args)
	return capsuleToPointer(a), capsuleToPointer(b), capsuleToPointer(c)
}

func parse4PointerFromArgs(args *C.PyObject) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	a, b, c, d := parse4ObjectFromArgs(args)
	return capsuleToPointer(a), capsuleToPointer(b), capsuleToPointer(c), capsuleToPointer(d)
}

func parse1ObjectFromArgs(args *C.PyObject) *C.PyObject {
	return C.PyTuple_GetItem(args, C.long(0))
}

func parse2ObjectFromArgs(args *C.PyObject) (*C.PyObject, *C.PyObject) {
	return C.PyTuple_GetItem(args, C.long(0)), C.PyTuple_GetItem(args, C.long(1))
}

func parse3ObjectFromArgs(args *C.PyObject) (*C.PyObject, *C.PyObject, *C.PyObject) {
	return C.PyTuple_GetItem(args, C.long(0)), C.PyTuple_GetItem(args, C.long(1)), C.PyTuple_GetItem(args, C.long(2))
}

func parse4ObjectFromArgs(args *C.PyObject) (*C.PyObject, *C.PyObject, *C.PyObject, *C.PyObject) {
	return C.PyTuple_GetItem(args, C.long(0)), C.PyTuple_GetItem(args, C.long(1)), C.PyTuple_GetItem(args, C.long(2)), C.PyTuple_GetItem(args, C.long(3))
}

func pointerToCapsule(ptr unsafe.Pointer) *C.PyObject {
	return C.influunt_pointerToCapsule(ptr)
}

func capsuleToPointer(o *C.PyObject) unsafe.Pointer {
	return C.PyCapsule_GetPointer(o, C.CString("influunt"))
}

func graphFromPointer(p unsafe.Pointer) *influunt.Graph {
	return pointer.Restore(p).(*influunt.Graph)
}

func modelFromPointer(p unsafe.Pointer) *influunt.Model {
	return pointer.Restore(p).(*influunt.Model)
}

func nodeToPointer(node influunt.Node) unsafe.Pointer {
	return pointer.Save(node)
}

func nodeFromPointer(p unsafe.Pointer) influunt.Node {
	node := pointer.Restore(p).(influunt.Node)
	return node
}
