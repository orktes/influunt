#define Py_LIMITED_API
#include <Python.h>

PyObject* influunt_NewGraph(PyObject* p0, PyObject* p1);
PyObject* influunt_ReadGraphFromFile(PyObject* p0, PyObject* p1);
PyObject* influunt_WriteGraphToFile(PyObject* p0, PyObject* p1);
PyObject* influunt_GraphAddOp(PyObject* p0, PyObject* p1);
PyObject* influunt_GraphNodeByName(PyObject* p0, PyObject* p1);

PyObject* influunt_NodeGetName(PyObject* p0, PyObject* p1);

PyObject* influunt_OpAdd(PyObject* p0, PyObject* p1);
PyObject* influunt_OpConst(PyObject* p0, PyObject* p1);
PyObject* influunt_OpPlaceholder(PyObject* p0, PyObject* p1);
PyObject* influunt_OpMap(PyObject* p0, PyObject* p1);
PyObject* influunt_OpParseJSON(PyObject* p0, PyObject* p1);
PyObject* influunt_OpCond(PyObject* p0, PyObject* p1);
PyObject* influunt_OpGetAttr(PyObject* p0, PyObject* p1);

PyObject* influunt_NewExecutor(PyObject* p0, PyObject* p1);
PyObject* influunt_ExecutorRun(PyObject* p0, PyObject* p1);

static PyMethodDef InfluuntMethods[] = {
    {"new_graph", influunt_NewGraph, METH_NOARGS, "Creates a new graph."},
    {"read_graph_from_file", influunt_ReadGraphFromFile, METH_O, "Read graph from file."},
    {"write_graph_to_file", influunt_WriteGraphToFile, METH_VARARGS, "Write graph to file."},
    {"graph_add_op", influunt_GraphAddOp, METH_VARARGS, "Add operation to graph."},
    {"graph_node_by_name", influunt_GraphNodeByName, METH_VARARGS, "Return node by name and index."},
    {"node_get_name", influunt_NodeGetName, METH_O, "Return node name."},
    {"new_executor", influunt_NewExecutor, METH_O, "Creates a new executor for a given graph."},
    {"executor_run", influunt_ExecutorRun, METH_VARARGS, "Execute contained graph with given input map and output filter"},
    {"op_add", influunt_OpAdd, METH_VARARGS, "Create an add operation between two nodes."},
    {"op_const", influunt_OpConst, METH_VARARGS, "Create new node with a constant value."},
    {"op_placeholder", influunt_OpPlaceholder, METH_VARARGS, "Create new placeholder node."},
    {"op_map", influunt_OpMap, METH_VARARGS, "Maps over a given list."},
    {"op_parsejson", influunt_OpParseJSON, METH_VARARGS, "Parses JSON."},
    {"op_getattr", influunt_OpGetAttr, METH_VARARGS, "Returns attribute of a dictionary."},
    {"op_cond", influunt_OpCond, METH_VARARGS, "Returs a if pred is \"true\" and b if pred is \"false\""},
    {NULL, NULL, 0, NULL}
};

static struct PyModuleDef influuntmodule = {  
   PyModuleDef_HEAD_INIT, "influunt_core", NULL, -1, InfluuntMethods
};

PyMODINIT_FUNC  
PyInit_influunt_core(void)  
{
    return PyModule_Create(&influuntmodule);
}