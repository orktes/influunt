package main

// #cgo pkg-config: python-3.6
// #include <Python.h>
//
// void influunt_pyRelease(PyObject *o)
// {
//	    Py_DECREF(o);
// }
//
// void influunt_pyRetain(PyObject *o)
// {
//	    Py_INCREF(o);
// }
//
// int influunt_isPyBool(PyObject *o)
// {
//	    return PyBool_Check(o);
// }
//
// int influunt_isPyBoolTrue(PyObject *o)
// {
//		return o == Py_True ? 1 : 0;
// }
//
// int influunt_isPyLong(PyObject *o)
// {
//		return PyLong_Check(o);
// }
//
// int influunt_isPyFloat(PyObject *o)
// {
//		return PyFloat_Check(o);
// }
//
// int influunt_isPyUnicode(PyObject *o)
// {
//		return PyUnicode_Check(o);
// }
//
// int influunt_isPyDict(PyObject *o)
// {
// 		return PyDict_Check(o);
// }
//
// int influunt_isPyList(PyObject *o)
// {
// 		return PyList_Check(o);
// }
//
// int influunt_isPyCapsule(PyObject *o)
// {
// 		return PyCapsule_CheckExact(o);
// }
//
// char* influunt_stringFromPyUnicode(PyObject *o)
// {
// 		return PyUnicode_AsUTF8AndSize(o, NULL);
// }
//
// void influunt_reprint(PyObject *obj) {
//    PyObject* repr = PyObject_Repr(obj);
//    PyObject* str = PyUnicode_AsEncodedString(repr, "utf-8", "~E~");
//    const char *bytes = PyBytes_AS_STRING(str);
//
//    printf("REPR: %s\n", bytes);
//
//    Py_XDECREF(repr);
//    Py_XDECREF(str);
// }
//
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/mattn/go-pointer"
)

type pyObject = C.PyObject

func pyPrint(pyObject *C.PyObject) {
	C.influunt_reprint(pyObject)
}

func pyRetain(pyObject *C.PyObject) {
	C.influunt_pyRetain(pyObject)
}

func pyRelease(pyObject *C.PyObject) {
	C.influunt_pyRelease(pyObject)
}

func convertGoTypeToPyObject(val interface{}) (*C.PyObject, error) {
	if val == nil {
		pyRetain(C.Py_None)
		return C.Py_None, nil
	}

	switch val := val.(type) {
	case reflect.Value:
		return convertGoTypeToPyObject(val.Interface())
	case int:
		return C.PyLong_FromLong(C.long(val)), nil
	case int32:
		return C.PyLong_FromLong(C.long(val)), nil
	case int64:
		return C.PyLong_FromLong(C.long(val)), nil
	case uint8:
		return C.PyLong_FromLong(C.long(val)), nil
	case uint:
		return C.PyLong_FromLong(C.long(val)), nil
	case uint16:
		return C.PyLong_FromLong(C.long(val)), nil
	case uint32:
		return C.PyLong_FromLong(C.long(val)), nil
	case uint64:
		return C.PyLong_FromUnsignedLong(C.ulong(val)), nil
	case float32:
		return C.PyFloat_FromDouble(C.double(val)), nil
	case float64:
		return C.PyFloat_FromDouble(C.double(val)), nil
	case string:
		return C.PyUnicode_FromString(C.CString(val)), nil
	case bool:
		if val {
			return C.Py_True, nil
		}

		return C.Py_False, nil
	case map[string]interface{}:
		dict := C.PyDict_New()

		for key, val := range val {
			pyKey, err := convertGoTypeToPyObject(key)
			if err != nil {
				return nil, err
			}

			pyVal, err := convertGoTypeToPyObject(val)
			if err != nil {
				return nil, err
			}

			C.PyDict_SetItem(dict, pyKey, pyVal)
			pyRelease(pyKey)
			pyRelease(pyVal)
		}

		return dict, nil
	case []interface{}:
		list := C.PyList_New(C.long(len(val)))

		for i, item := range val {
			pyItem, err := convertGoTypeToPyObject(item)
			if err != nil {
				return nil, err
			}

			C.PyList_SetItem(list, C.long(i), pyItem)
		}

		return list, nil
	case unsafe.Pointer:
		return pointerToCapsule(val), nil
	default:
		typ := reflect.TypeOf(val)
		switch typ.Kind() {
		case reflect.Ptr:
			return convertGoTypeToPyObject(reflect.ValueOf(val).Elem())
		case reflect.Slice, reflect.Array:
			val := reflect.ValueOf(val)
			length := val.Len()
			list := C.PyList_New(C.long(length))

			for i := 0; i < length; i++ {
				item := val.Index(i)
				pyItem, err := convertGoTypeToPyObject(item)
				if err != nil {
					return nil, err
				}

				C.PyList_SetItem(list, C.long(i), pyItem)
			}

			return list, nil
		case reflect.Map:
			dict := C.PyDict_New()
			val := reflect.ValueOf(val)
			for _, key := range val.MapKeys() {
				val := val.MapIndex(key)

				pyKey, err := convertGoTypeToPyObject(key)
				if err != nil {
					return nil, err
				}

				pyVal, err := convertGoTypeToPyObject(val)
				if err != nil {
					return nil, err
				}

				C.PyDict_SetItem(dict, pyKey, pyVal)
			}

			return dict, nil
		}
	}

	return nil, fmt.Errorf("unsupported type %T", val)
}

func convertPyObjectToInterface(pyObject *C.PyObject) (interface{}, error) {
	if pyObject == nil {
		return nil, nil
	}

	if int(C.influunt_isPyBool(pyObject)) == 1 {
		return C.influunt_isPyBoolTrue(pyObject) == 1, nil
	} else if int(C.influunt_isPyLong(pyObject)) == 1 {
		long := C.PyLong_AsLong(pyObject)
		return int(long), nil
	} else if int(C.influunt_isPyFloat(pyObject)) == 1 {
		double := C.PyFloat_AsDouble(pyObject)
		return float64(double), nil
	} else if int(C.influunt_isPyUnicode(pyObject)) == 1 {
		return C.GoString(C.influunt_stringFromPyUnicode(pyObject)), nil
	} else if int(C.influunt_isPyCapsule(pyObject)) == 1 {
		return pointer.Restore(capsuleToPointer(pyObject)), nil
	} else if int(C.influunt_isPyList(pyObject)) == 1 {
		length := int(C.PyList_Size(pyObject))
		if length == 0 {
			return []interface{}{}, nil
		}

		var arr reflect.Value

		for i := 0; i < length; i++ {
			item := C.PyList_GetItem(pyObject, C.long(i))
			val, err := convertPyObjectToInterface(item)
			if err != nil {
				return nil, err
			}

			if i == 0 {
				arr = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(val)), length, length)
			}

			v := arr.Index(i)
			v.Set(reflect.ValueOf(val))
		}

		return arr.Interface(), nil
	} else if int(C.influunt_isPyDict(pyObject)) == 1 {
		m := map[string]interface{}{}

		keys := C.PyDict_Keys(pyObject)
		length := int(C.PyList_Size(keys))

		for i := 0; i < length; i++ {
			key := C.PyList_GetItem(keys, C.long(i))
			keyVal, err := convertPyObjectToInterface(key)
			if err != nil {
				return nil, err
			}

			if strKey, ok := keyVal.(string); ok {
				val := C.PyDict_GetItem(pyObject, key)
				valVal, err := convertPyObjectToInterface(val)
				if err != nil {
					return nil, err
				}

				m[strKey] = valVal
			} else {
				return nil, errors.New("only string based keys are supported currently")
			}
		}
		return m, nil
	}
	return nil, errors.New("python type not supported")
}
