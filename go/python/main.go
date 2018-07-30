package main

// #cgo pkg-config: python-3.6
// #include <Python.h>
import "C"
import "github.com/mattn/go-pointer"

// influunt_destructCapsule destructs a capsule
//export influunt_destructCapsule
func influunt_destructCapsule(c *C.PyObject) {
	pointer.Unref(capsuleToPointer(c))
}

func main() {
	// HERE be dragons. This is build a shared c library
}
