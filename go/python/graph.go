package main

import "C"
import (
	"github.com/mattn/go-pointer"
	"github.com/orktes/influunt/go"
)

// influunt_NewGraph returns a new influunt graph as void pointer
//export influunt_NewGraph
func influunt_NewGraph(self *pyObject, args *pyObject) *pyObject {
	return pointerToCapsule(pointer.Save(influunt.NewGraph()))
}
